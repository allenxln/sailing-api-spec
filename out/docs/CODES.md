# Error code table

> Generated from `codes.yaml`. Do not edit directly.

## Ranges

| Range | Start | End | Description |
|-------|-------|-----|-------------|
| `common` | 0 | 9999 | Cross-service common codes |
| `tutor_core` | 21000 | 21999 | tutor_core business codes |
| `words_core` | 20000 | 20999 | words_core business codes |

## Codes

| Code | Name | 中文 | English | HTTP | Range |
|------|------|------|---------|------|-------|
| 0 | `Success` | 成功 | ok | 200 | `common` |
| 1001 | `LoginStatusExpired` | 登录已过期 | auth status expired | 401 | `common` |
| 1003 | `PermissionDenied` | 权限不足 | permission denied | 403 | `common` |
| 4001 | `BadRequest` | 请求参数错误 | bad request | 400 | `common` |
| 5001 | `InternalServerError` | 服务器内部错误 | internal server error | 500 | `common` |
| 20001 | `PackageNotFound` | 词包不存在 | package not found | 404 | `words_core` |
| 21001 | `KnowledgeNameConflict` | 知识点名称已存在 | knowledge name already exists | 409 | `tutor_core` |
| 21002 | `VerifyCodeInvalid` | 验证码错误或已过期 | verification code invalid or expired | 400 | `tutor_core` |
| 21003 | `EmailAlreadyRegistered` | 邮箱已被注册 | email already registered | 409 | `tutor_core` |
| 21004 | `LoginFailed` | 邮箱或密码错误 | email or password wrong | 401 | `tutor_core` |
| 21005 | `VerifyCodeSendLimited` | 验证码发送过于频繁，请稍后再试 | verification code sent too frequently | 429 | `tutor_core` |
| 21006 | `TooManyLoginAttempts` | 尝试次数过多，请稍后再试 | too many login attempts | 429 | `tutor_core` |
| 21007 | `TreeNodeInvalid` | 树节点操作非法（移到自身子树下等） | tree node operation invalid | 400 | `tutor_core` |
| 21008 | `TreeNodeNameConflict` | 同级下已存在同名节点 | sibling node name conflict | 409 | `tutor_core` |
| 21009 | `QuestionNotFound` | 题目不存在 | question not found | 404 | `tutor_core` |
| 21010 | `VerifyTokenInvalid` | 验证链接无效或已过期 | verification token invalid or expired | 400 | `tutor_core` |
