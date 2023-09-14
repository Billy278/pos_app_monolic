package responses

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidPayload     = "invalid payload request"
	InvalidQuery       = "invalid query request"
	InternalServer     = "internal server error"
	SomethingWentWrong = "something went wrong"
	Unauthorized       = "unauthorized request"
	Success            = "Success"
	NotFound           = "NOT FOUND"
)

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
