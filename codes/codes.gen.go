// Code generated from codes.yaml. DO NOT EDIT.
// Edit codes.yaml then run `make gen`.

package codes

import "github.com/zeromicro/x/errors"

var (
	// Success: 成功 (ok)
	Success = errors.New(0, "ok")
	// LoginStatusExpired: 登录已过期 (auth status expired)
	LoginStatusExpired = errors.New(1001, "auth status expired")
	// PermissionDenied: 权限不足 (permission denied)
	PermissionDenied = errors.New(1003, "permission denied")
	// BadRequest: 请求参数错误 (bad request)
	BadRequest = errors.New(4001, "bad request")
	// InternalServerError: 服务器内部错误 (internal server error)
	InternalServerError = errors.New(5001, "internal server error")
	// PackageNotFound: 词包不存在 (package not found)
	PackageNotFound = errors.New(20001, "package not found")
	// KnowledgeNameConflict: 知识点名称已存在 (knowledge name already exists)
	KnowledgeNameConflict = errors.New(21001, "knowledge name already exists")
	// VerifyCodeInvalid: 验证码错误或已过期 (verification code invalid or expired)
	VerifyCodeInvalid = errors.New(21002, "verification code invalid or expired")
	// EmailAlreadyRegistered: 邮箱已被注册 (email already registered)
	EmailAlreadyRegistered = errors.New(21003, "email already registered")
	// LoginFailed: 邮箱或密码错误 (email or password wrong)
	LoginFailed = errors.New(21004, "email or password wrong")
	// VerifyCodeSendLimited: 验证码发送过于频繁，请稍后再试 (verification code sent too frequently)
	VerifyCodeSendLimited = errors.New(21005, "verification code sent too frequently")
	// TooManyLoginAttempts: 尝试次数过多，请稍后再试 (too many login attempts)
	TooManyLoginAttempts = errors.New(21006, "too many login attempts")
	// TreeNodeInvalid: 树节点操作非法（移到自身子树下等） (tree node operation invalid)
	TreeNodeInvalid = errors.New(21007, "tree node operation invalid")
	// TreeNodeNameConflict: 同级下已存在同名节点 (sibling node name conflict)
	TreeNodeNameConflict = errors.New(21008, "sibling node name conflict")
	// QuestionNotFound: 题目不存在 (question not found)
	QuestionNotFound = errors.New(21009, "question not found")
	// VerifyTokenInvalid: 验证链接无效或已过期 (verification token invalid or expired)
	VerifyTokenInvalid = errors.New(21010, "verification token invalid or expired")
	// TreeNotFound: 知识树不存在 (tree not found)
	TreeNotFound = errors.New(21011, "tree not found")
	// TreeAlreadyExists: 同维度知识树已存在 (tree with same dimensions already exists)
	TreeAlreadyExists = errors.New(21012, "tree with same dimensions already exists")
	// KnowledgeAlreadyMounted: 该知识点已挂载在此节点 (knowledge point already mounted on this node)
	KnowledgeAlreadyMounted = errors.New(21013, "knowledge point already mounted on this node")
)

// HTTPStatus maps a code to its recommended HTTP status.
func HTTPStatus(code int) int {
	switch code {
	case 0:
		return 200
	case 1001:
		return 401
	case 1003:
		return 403
	case 4001:
		return 400
	case 5001:
		return 500
	case 20001:
		return 404
	case 21001:
		return 409
	case 21002:
		return 400
	case 21003:
		return 409
	case 21004:
		return 401
	case 21005:
		return 429
	case 21006:
		return 429
	case 21007:
		return 400
	case 21008:
		return 409
	case 21009:
		return 404
	case 21010:
		return 400
	case 21011:
		return 404
	case 21012:
		return 409
	case 21013:
		return 409
	default:
		return 500
	}
}
