# sailing-api-spec

Single source of truth for error codes (and future API schemas) shared by
Sailing backend services (`words_core`, future `user_core`, ...) and
clients (`sailing_words` Flutter app).

## Why

Without a contract repo, every new error code has to be hand-added in at
least three places — the Go backend, the Dart client, and whatever doc the
PM / QA are reading — and they inevitably drift. A contract repo makes the
yaml the only place to edit, and generates Go, Dart, and Markdown from it.

```
codes.yaml                       ← humans edit this
    │
    │  make gen
    ▼
codes/codes.gen.go               ← Go backend imports this
out/dart/lib/error_code.dart     ← Flutter app imports this
out/docs/CODES.md                ← PM / QA / new hires read this
```

## Adding a new error code

1. Edit `codes.yaml`, add an entry under `codes:`. Pick a code inside the
   right range (see `ranges:` at the top of the yaml).
2. Run `make gen` — this validates the yaml (no duplicate codes, no out-of-range
   numbers, every code has zh + en + http) and regenerates the three outputs.
3. Commit everything (the yaml **and** the regenerated `out/` tree).
4. Open a PR. CI runs `make check` and `make test-go`; if you forgot to run
   `make gen`, or the yaml was malformed, the PR turns red.
5. Once merged to `main`, tag a release: `git tag vX.Y.Z && git push --tags`.
   A GitHub Release is published automatically.

## Consuming the contract

### Go backend

```bash
go get github.com/allenxln/sailing-api-spec/codes@vX.Y.Z
```

```go
import "github.com/allenxln/sailing-api-spec/codes"

func (l *FooLogic) Foo(req *FooReq) (*FooResp, error) {
    if missing {
        return nil, codes.PackageNotFound  // typed, no magic 20001
    }
    ...
}
```

The codes are `errors.CodeMsg` values from `github.com/zeromicro/x/errors`,
so they work with go-zero's existing envelope/response plumbing out of the box.

### Flutter client

In `pubspec.yaml`:

```yaml
dependencies:
  sailing_contract:
    git:
      url: git@github.com:allenxln/sailing-api-spec.git
      path: out/dart
      ref: vX.Y.Z
```

```dart
import 'package:sailing_contract/error_code.dart';

if (code == ErrorCode.loginStatusExpired) {
  // no magic 1001
  showReloginDialog();
} else {
  showError(ErrorCode.messageZh(code));
}
```

## Repo layout

```
.
├── codes.yaml                   # edit this
├── gen/
│   ├── main.go                  # generator (validate + render)
│   └── templates/
│       ├── go.tmpl
│       ├── dart.tmpl
│       └── md.tmpl
├── codes/                       # committed Go package (go get target)
│   └── codes.gen.go
├── out/                         # committed non-Go artifacts
│   ├── dart/lib/error_code.dart
│   └── docs/CODES.md
├── Makefile
└── .github/workflows/
    ├── validate.yml             # PR gate: yaml ↔ out/ in sync + Go compiles
    └── release.yml              # tag push → GitHub Release
```

## Code range policy

Every code belongs to a `range`. Adding a new service? Reserve a 1000-wide
block in `ranges:` and stick to it. Current allocation:

| Range | Span | Owner |
|-------|------|-------|
| `common` | 0–9999 | cross-service (auth, bad request, server error) |
| `words_core` | 20000–20999 | words_core |

Collisions, name dupes, out-of-range codes, and missing zh/en/http are all
rejected by `make gen`. Don't chase around — let the generator tell you.

## Versioning

Semver-ish:

- **Patch** (`v1.0.1`): fix a msg typo, adjust http status on an existing code.
- **Minor** (`v1.1.0`): add new codes. Existing codes unchanged.
- **Major** (`v2.0.0`): change or delete a previously-published code number.
  Avoid this — every client in production needs to redeploy in lockstep.

## Why generate into `out/` instead of just a Go module?

Because clients are not just Go. If this were Go-only we could just write
hand-maintained Go consts. Keeping yaml as the source and generating into
multiple target languages is the whole point.
