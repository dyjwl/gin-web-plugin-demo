package resp

import "errors"

var msgcode = map[int]string{
	Success: "ok",

	ErrSignatureInvalid:  "无法识别的请求头",
	ErrBadRequest:        "请求参数错误",
	ErrInvalidAuthHeader: "非法的请求头",
	ErrPageNotFound:      "找不到页面",
	ErrPermissionDenied:  "无操作权限",
	ErrBind:              "参数异常",
	ErrValidation:        "请求参数校验异常",
	ErrInternalServer:    "服务器内部错误",
}

func GetAPIMsgByCode(code int) string {
	if msg, ok := msgcode[code]; ok {
		return msg
	}
	return "unknown"
}

func GetAPIMsgErr(code int) error {
	return errors.New(GetAPIMsgByCode(code))
}
