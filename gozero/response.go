// Package gozero provides the {code, msg, data} JSON envelope renderer used
// by every go-zero handler in the Sailing backend. It lives in the same
// repository as the error code contract (../codes/) so a service only ever
// needs to depend on one module.
package gozero

import (
	"net/http"

	stderrors "errors"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

// Body is the canonical response envelope. Every HTTP response from every
// Sailing service is shaped like this.
type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response writes {code, msg, data}.
//
// - err == nil          → code=0, msg="ok", data=resp.
// - err is/wraps a *errors.CodeMsg → code/msg propagated verbatim, no data.
// - any other err       → code=5001, msg="internal server error".
//
// The type-assertion path (errors.As) is what lets logic layers return
// typed codes defined in any package — most commonly the generated
// constants in ../codes — without this function needing to know about them.
func Response(w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		httpx.OkJson(w, Body{Code: 0, Msg: "ok", Data: resp})
		return
	}
	var cm *errors.CodeMsg
	if stderrors.As(err, &cm) {
		httpx.OkJson(w, Body{Code: cm.Code, Msg: cm.Msg})
		return
	}
	httpx.OkJson(w, Body{Code: 5001, Msg: "internal server error"})
}

// BadRequest is a convenience that handlers call when request parsing
// (httpx.Parse) fails. Keeps handler templates small.
func BadRequest(w http.ResponseWriter, err error) {
	httpx.OkJson(w, Body{Code: 4001, Msg: "bad request"})
}
