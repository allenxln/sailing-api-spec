# 本仓的恢复说明

> 本文件由 2026-08-16 的恢复操作添加，不属于原始仓库内容。

原仓 `allenxln/sailing-api-spec` 于 2026-07-07（最后一个版本 v0.10.0 发布）之后的某时被误删。
本仓由 Go module 代理缓存（goproxy.cn 与 proxy.golang.org）中留存的各版本模块 zip 重建：

- **文件内容**：每个 tag 的文件树与原版逐字节一致。v0.6.0 已通过 `words_core` go.sum 中的
  模块哈希验证（`h1:3D1JVh736dOHtgm1ckIZSuVGFYagl4Cj8ytD8hR/wDY=`，与 sum.golang.org 记录一致）。
- **git 历史**：无法恢复。历史被压缩为每个版本一个提交，提交时间取自代理记录的版本时间。
- **原始 commit sha**（已知部分，来自代理 Origin 记录与 sailing_words pubspec.lock）：

  | tag | 原始 commit |
  |---|---|
  | v0.4.0 | `ab56764c11c59d4c0cb4d24c00249040eeaedd27` |
  | v0.6.0 | `953ccf202b043694d2147c4215cc2a58521985d9` |
  | v0.9.0 | `810c99ae1a437dc1acb6f0b5648a8c8c4133abd0` |
  | v0.9.1 | `2f313e9bb9d76a1c133e99465ac61987a046f480` |
  | v0.10.0 | `3256503578ee9213ba1fc610fda22548a8d81914` |

恢复方式：从 `proxy.golang.org` 与 `goproxy.cn` 下载全部 11 个版本的模块 zip，
按版本时间顺序重建提交并打 tag。依赖本仓的消费方（`words_core` go.mod、
`sailing_words` pubspec.yaml 均钉在 v0.6.0）无需任何改动即可继续使用。
