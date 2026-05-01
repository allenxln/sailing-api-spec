# Error code table

> Generated from `codes.yaml`. Do not edit directly.

## Ranges

| Range | Start | End | Description |
|-------|-------|-----|-------------|
| `common` | 0 | 9999 | Cross-service common codes |
| `words_core` | 20000 | 20999 | words_core business codes |

## Codes

| Code | Name | 中文 | English | HTTP | Range |
|------|------|------|---------|------|-------|
| 0 | `Success` | 成功 | ok | 200 | `common` |
| 1001 | `LoginStatusExpired` | 登录已过期 | auth status expired | 401 | `common` |
| 4001 | `BadRequest` | 请求参数错误 | bad request | 400 | `common` |
| 5001 | `InternalServerError` | 服务器内部错误 | internal server error | 500 | `common` |
| 20001 | `PackageNotFound` | 词包不存在 | package not found | 404 | `words_core` |
