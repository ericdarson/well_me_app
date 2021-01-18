package response

type CheckSessionAdminOutputSchema struct {
	SystemDate         string                  `json:"system_date"`
	DetailCheckSession DetailCheckAdminSession `json:"session"`
}

type DetailCheckAdminSession struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Token    string `json:"new_token"`
}
