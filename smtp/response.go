package smtp

// Interfaceでもいいかもしれない
// サーバ→クライアントの返り値
type Response struct {
	code       int
	followNext bool
	text       string
}

func (r *Response) Code() int {
	return r.code
}

func (r *Response) Text() string {
	return r.text
}

// TODO: Validate code
func NewResponse(code int, text string) *Response {
	return &Response{
		code: code,
		text: text,
	}
}
