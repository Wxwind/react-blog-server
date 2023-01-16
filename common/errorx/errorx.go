package errorx

const (
	DATABASE_MYSQL_INTERNAL_ERROR     = 1001
	DATABASE_MYSQL_NOT_EXISTS_ARTICLE = 1002
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) ToData() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
		Data: nil,
	}
}
