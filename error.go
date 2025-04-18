package ginresponse

type ErrorModel struct {
	Message string `json:"message"`
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}