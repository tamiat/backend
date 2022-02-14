package response

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}