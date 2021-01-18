package response

type CheckSessionOutputSchema struct {
	SystemDate         string             `json:"system_date"`
	DetailCheckSession DetailCheckSession `json:"session"`
}

type DetailCheckSession struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Token    string `json:"new_token"`
}
