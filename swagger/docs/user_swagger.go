package docs

import "github.com/dyjwl/gin-web-plugin-demo/internal/store/model"

// swagger:route POST /app/api/v1/user/register users createUserRequest
//
// 用户注册.
//
// 根据参数注册用户.
//
//     Security:
//       api_key:
//
//     Responses:
//       default: errResponse
//       200: okResponse

// swagger:parameters createUserRequest
type userRequestParamsWrapper struct {
	// User information.
	// in:body
	Body model.User
}

// Return nil json object.
// swagger:response okResponse
type okResponseWrapper struct{}

// swagger:response errResponse
type errResponseWrapper struct{}
