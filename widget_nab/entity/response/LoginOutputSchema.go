package response

type LoginOutputSchema struct {
	SystemDate  string      `json:"system_date"`
	DetailLogin DetailLogin `json:"detail_login"`
}

type DetailLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
