package response

type LoginAdminOutputSchema struct {
	SystemDate       string           `json:"system_date"`
	DetailLoginAdmin DetailLoginAdmin `json:"detail_login"`
}

type DetailLoginAdmin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
