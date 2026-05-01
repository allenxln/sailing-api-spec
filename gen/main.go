package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v3"
)

type Range struct {
	Start int    `yaml:"start"`
	End   int    `yaml:"end"`
	Desc  string `yaml:"desc"`
}

type Code struct {
	Code  int    `yaml:"code"`
	Name  string `yaml:"name"`
	Range string `yaml:"range"`
	MsgZh string `yaml:"msg_zh"`
	MsgEn string `yaml:"msg_en"`
	HTTP  int    `yaml:"http"`
}

type Spec struct {
	Version int              `yaml:"version"`
	Ranges  map[string]Range `yaml:"ranges"`
	Codes   []Code           `yaml:"codes"`
}

//go:embed templates/go.tmpl
var goTmpl string

//go:embed templates/dart.tmpl
var dartTmpl string

//go:embed templates/md.tmpl
var mdTmpl string

func main() {
	spec, err := parse("codes.yaml")
	if err != nil {
		die("parse: %v", err)
	}
	if err := validate(spec); err != nil {
		die("validate: %v", err)
	}
	// Stable ordering makes generated diffs reviewable.
	sort.Slice(spec.Codes, func(i, j int) bool { return spec.Codes[i].Code < spec.Codes[j].Code })

	mustRender(goTmpl, "out/go/codes/codes.gen.go", spec, nil)
	mustRender(dartTmpl, "out/dart/lib/error_code.dart", spec, template.FuncMap{
		"lowerCamel": lowerCamel,
	})
	mustRender(mdTmpl, "out/docs/CODES.md", spec, template.FuncMap{
		"rangeDesc": func(name string) string { return spec.Ranges[name].Desc },
	})
	fmt.Printf("OK: %d codes → 3 outputs\n", len(spec.Codes))
}

func parse(path string) (Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Spec{}, err
	}
	var s Spec
	if err := yaml.Unmarshal(data, &s); err != nil {
		return Spec{}, err
	}
	return s, nil
}

func validate(s Spec) error {
	if s.Version == 0 {
		return fmt.Errorf("missing version")
	}
	if len(s.Ranges) == 0 {
		return fmt.Errorf("no ranges declared")
	}
	seenCode := map[int]string{}
	seenName := map[string]bool{}
	for _, c := range s.Codes {
		r, ok := s.Ranges[c.Range]
		if !ok {
			return fmt.Errorf("code %d (%s): unknown range %q", c.Code, c.Name, c.Range)
		}
		if c.Code < r.Start || c.Code > r.End {
			return fmt.Errorf("code %d (%s) out of range %s [%d-%d]",
				c.Code, c.Name, c.Range, r.Start, r.End)
		}
		if prev, dup := seenCode[c.Code]; dup {
			return fmt.Errorf("duplicate code %d: %s vs %s", c.Code, prev, c.Name)
		}
		seenCode[c.Code] = c.Name
		if seenName[c.Name] {
			return fmt.Errorf("duplicate name %s", c.Name)
		}
		seenName[c.Name] = true
		if c.MsgZh == "" || c.MsgEn == "" {
			return fmt.Errorf("code %d (%s): missing msg_zh or msg_en", c.Code, c.Name)
		}
		if c.HTTP == 0 {
			return fmt.Errorf("code %d (%s): missing http status", c.Code, c.Name)
		}
	}
	return nil
}

func mustRender(tmpl, outPath string, s Spec, funcs template.FuncMap) {
	t := template.New(filepath.Base(outPath))
	if funcs != nil {
		t = t.Funcs(funcs)
	}
	t = template.Must(t.Parse(tmpl))

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		die("mkdir %s: %v", outPath, err)
	}
	f, err := os.Create(outPath)
	if err != nil {
		die("create %s: %v", outPath, err)
	}
	defer f.Close()
	if err := t.Execute(f, s); err != nil {
		die("render %s: %v", outPath, err)
	}
}

// lowerCamel converts "PackageNotFound" -> "packageNotFound" (for Dart constants).
func lowerCamel(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	os.Exit(1)
}

var _ = strings.ToLower // keep strings imported for future use
