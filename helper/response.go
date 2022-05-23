package helper

type responJSON struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Response(status bool, code int, message string, body interface{}) *responJSON {
	return &responJSON{
		Status:  status,
		Code:    code,
		Message: message,
		Body:    body,
	}
}
