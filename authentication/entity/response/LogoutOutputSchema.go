package response

type LogoutOutputSchema struct {
	SystemDate   string       `json:"system_date"`
	DetailLogout DetailLogout `json:"detail_logout"`
}

type DetailLogout struct {
	Message string `json:"message"`
}
