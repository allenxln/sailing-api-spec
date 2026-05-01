// Code generated from codes.yaml. DO NOT EDIT.
// Edit codes.yaml then run `make gen`.

package codes

import "github.com/zeromicro/x/errors"

var (
	// Success: 成功 (ok)
	Success = errors.New(0, "ok")
	// LoginStatusExpired: 登录已过期 (auth status expired)
	LoginStatusExpired = errors.New(1001, "auth status expired")
	// BadRequest: 请求参数错误 (bad request)
	BadRequest = errors.New(4001, "bad request")
	// InternalServerError: 服务器内部错误 (internal server error)
	InternalServerError = errors.New(5001, "internal server error")
	// PackageNotFound: 词包不存在 (package not found)
	PackageNotFound = errors.New(20001, "package not found")
)

// HTTPStatus maps a code to its recommended HTTP status.
func HTTPStatus(code int) int {
	switch code {
	case 0:
		return 200
	case 1001:
		return 401
	case 4001:
		return 400
	case 5001:
		return 500
	case 20001:
		return 404
	default:
		return 500
	}
}
