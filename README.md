# sailing-api-spec

Sailing 所有后端服务（`words_core`、未来的 `user_core` …）和前端（`sailing_words` Flutter App）**共享的错误码契约**。错误码在这里统一定义、自动生成各端代码，不再靠约定俗成或手抄对齐。

## 为什么要有这个仓库

没有契约仓时，加一个新错误码要在至少三个地方同步：Go 后端常量、Dart 客户端常量、产品/测试读的文档。三者不可避免地漂移——后端改了 msg，前端不知道；前端加了错误 UI 的魔法数字 `20001`，后端重构时没人提醒。

契约仓把 yaml 当作唯一真相，Go / Dart / Markdown 都是它的机器生成物。

```
codes.yaml                       ← 人改这里
    │
    │  make gen
    ▼
codes/codes.gen.go               ← Go 后端 import 这个
gozero/response.go               ← Go handler 调 gozero.Response(w, resp, err)
out/dart/lib/error_code.dart     ← Flutter 客户端 import 这个
out/docs/CODES.md                ← PM / 测试 / 新人看这个
```

## 加一个新错误码

1. 编辑 `codes.yaml`，在 `codes:` 下加一条。code 要落在对应 `range:` 区间内（见 yaml 顶部）。
2. 跑 `make gen`：校验 yaml（撞码、越界、缺 zh/en/http 都会在这一步报错），同时重新生成三份产物。
3. commit：yaml **和**重新生成的 `codes/` + `out/` 一起入库。
4. 开 PR。CI 会跑 `make check` 和 `make test-go`。忘跑 `make gen` 或 yaml 写错 → PR 红。
5. 合到 `main` 后，打 tag 发版：`git tag vX.Y.Z && git push --tags`。CI 会自动发一个 GitHub Release。

## 下游怎么消费

### Go 后端

```bash
go get github.com/allenxln/sailing-api-spec@vX.Y.Z
```

业务 logic 里：

```go
import (
    "github.com/allenxln/sailing-api-spec/codes"
    "github.com/allenxln/sailing-api-spec/gozero"
)

func (l *FooLogic) Foo(req *FooReq) (*FooResp, error) {
    if missing {
        return nil, codes.PackageNotFound  // 不再是 return nil, fmt.Errorf("...")
    }
    ...
}
```

handler 模板里（或者手写）：

```go
resp, err := l.Foo(&req)
gozero.Response(w, resp, err)  // 统一 {code, msg, data} envelope
```

`codes.XXX` 都是 `*errors.CodeMsg`（来自 `github.com/zeromicro/x/errors`），可以被 `gozero.Response` 透传成 `{"code":20001,"msg":"..."}`。任何没登记的 error 会降级成 `5001 internal server error`。

### Flutter 客户端

`pubspec.yaml`：

```yaml
dependencies:
  sailing_contract:
    git:
      url: https://github.com/allenxln/sailing-api-spec.git
      path: out/dart
      ref: vX.Y.Z
```

Dart 代码里：

```dart
import 'package:sailing_contract/error_code.dart';

if (code == ErrorCode.loginStatusExpired) {
  // 不再是魔法数字 1001
  showReloginDialog();
} else {
  showError(ErrorCode.messageZh(code));
}
```

## 仓库结构

```
.
├── codes.yaml                   # 人改的唯一真相
├── codes/
│   └── codes.gen.go             # Go 消费入口（go get target）
├── gozero/
│   └── response.go              # {code,msg,data} 渲染器
├── gen/
│   ├── main.go                  # 生成器 + 校验器
│   └── templates/
│       ├── go.tmpl
│       ├── dart.tmpl
│       ├── pubspec.tmpl
│       └── md.tmpl
├── out/                         # 入库的非 Go 产物
│   ├── dart/
│   │   ├── pubspec.yaml         # Flutter 把这当作一个 Dart package
│   │   └── lib/error_code.dart
│   └── docs/CODES.md            # 人读的码表
├── Makefile                     # gen / check / test-go / clean
└── .github/workflows/
    ├── validate.yml             # PR gate: yaml ↔ 产物一致 + Go 能编译
    └── release.yml              # tag push → GitHub Release
```

说明：

- **Go 产物在仓库根 (`codes/`)，其他产物在 `out/`**。不对称是刻意的：Go module 要求入口在 module root 附近，否则消费方要多加一层子 module + 多模块 tag；Dart 和 Markdown 没这个约束，放哪都行。
- **产物入库**。这样消费方 `go get` / `pub get` 直接能用，不需要本地装生成器。
- **所有 `*.gen.go` / `error_code.dart` / `CODES.md` 都带 `DO NOT EDIT` 标记**。想改要改 yaml，不是改产物。CI 的 `make check` 会拿 `git diff` 卡住"改了 yaml 但忘跑 gen"的 PR。

## 错误码区间约定

每条 code 都属于一个 `range`。新接一个服务就在 `ranges:` 里申请一块 1000 号段，code 写在区间内。当前分配：

| Range | 号段 | 归属 |
|-------|------|------|
| `common` | 0–9999 | 跨服务通用（登录、参数错、系统错等） |
| `words_core` | 20000–20999 | words_core 业务码 |

撞码、撞名、越界、缺 zh/en/http —— 全部由 `make gen` 在生成前拦下，不用靠人眼 review 接住。

## 版本策略（Semver 近似）

- **Patch**（`v1.0.1`）：修 msg typo、微调已有码的 http 状态。
- **Minor**（`v1.1.0`）：新增错误码。已有码**不动**。
- **Major**（`v2.0.0`）：改或删一个已发布的 code 数字。**尽量别用**——每个线上客户端都要同步升级，协议破坏级别的变更。

消费方 pin 具体 tag（`@v0.5.0` / `ref: v0.5.0`），不要追 `main`。要升级靠 `go get .../latest` 或改一行 `ref:`，或者配 Renovate 自动开升级 PR。

## 为什么生成到 `out/` / `codes/` 而不是只发一个 Go module

因为消费方不止 Go。如果只有 Go，直接手写 const 也能活。yaml → 多语言产物这个设计的全部意义，就是**让 Flutter / 未来的 Web / 未来的文档站共享同一份真相**。

## 灵感来源 / 设计参考

这个仓库不是凭空拍脑袋的——它是几个业界成熟模式的组合，每一块都有公开的大厂实践可以对标。列出来方便未来接手的人理解"为什么这么设计"。

### "错误码用 schema 作为 SSOT + 多语言生成"

- **[Google APIs](https://github.com/googleapis/googleapis)** `google/rpc/code.proto` —— Google Cloud 全球 SDK 的错误码共享源
- **[gRPC Status Codes](https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto)** —— 跨语言标准错误码
- **[Kubernetes API errors](https://github.com/kubernetes/apimachinery/blob/master/pkg/api/errors/errors.go)**
- **[Stripe OpenAPI](https://github.com/stripe/openapi)** —— Stripe 11 种语言 SDK 都从这一份 spec 生成
- **[Google AIP-193 Errors](https://google.aip.dev/193)** —— Google 内部 API 错误设计规范

### Contract-first / Schema-first design

- Martin Fowler —— [Consumer-Driven Contracts](https://martinfowler.com/articles/consumerDrivenContracts.html)
- Shopify Engineering —— [Evolving API Pipelines at Shopify](https://shopify.engineering/evolving-api-pipelines)
- **[Buf](https://buf.build/)** —— 专门做 proto schema 仓库托管 + 多语言生成的商业方案

### goctl `--remote` 共享模板

- go-zero 官方文档 —— [goctl 模板自定义](https://go-zero.dev/docs/tasks/cli/template)
- **[`zeromicro/go-zero-template`](https://github.com/zeromicro/go-zero-template)** —— 官方样板，本仓的 `api/handler.tpl` 遵循它的目录约定

### 老牌 schema-first 方案（思路一致，更重）

- **[Protocol Buffers](https://protobuf.dev/)**
- **[Apache Thrift](https://thrift.apache.org/)**
- **[Apache Avro](https://avro.apache.org/)**

### 业内定位

这套方案在业界通常叫 **"API Contract Repository"** 或 **"Schema-First Code Generation"**。搜关键词：

- `api contract repository`
- `schema-first api design`
- `error code registry`
- `single source of truth api codes`
- `cross-language code generation from schema`

**诚实说明**：没有哪个开源项目可以直接照抄 —— 因为每家团队的错误码表、envelope 格式、技术栈（Go-zero + Flutter 的组合在中国团队里常见但英文圈偏主流的是 gRPC + TS/Python）都不一样。这个仓库是用上面的积木，拼给 Sailing 技术栈的一个最小实例。
