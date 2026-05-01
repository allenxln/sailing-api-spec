.PHONY: gen check test-go fmt clean

# Run the generator against codes.yaml. Produces:
#   out/go/codes/codes.gen.go
#   out/dart/lib/error_code.dart
#   out/docs/CODES.md
gen:
	go run ./gen

# CI gate: yaml → generated outputs must be in sync. Forces anyone editing
# codes.yaml to also commit the regenerated out/ tree.
check: gen
	@git diff --exit-code -- out/ \
	  || (echo "ERROR: out/ is out of sync with codes.yaml. Run 'make gen' and commit."; exit 1)

# Verify the generated Go package compiles. Catches template drift
# (e.g. syntax error in a template) before downstream services pull the tag.
test-go: gen
	go build ./codes/...

fmt:
	gofmt -w gen/

clean:
	rm -rf out/
