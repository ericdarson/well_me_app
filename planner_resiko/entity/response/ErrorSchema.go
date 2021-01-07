package response

type ErrorSchema struct {
	ErrorCode    string       `json:"error_code"`
	ErrorMessage ErrorMessage `json:"error_message"`
}

type ErrorMessage struct {
	English    string `json:"english"`
	Indonesian string `json:"indonesian"`
}
