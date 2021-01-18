package response

type LogoutResponse struct {
	ErrorSchema  ErrorSchema        `json:"error_schema"`
	OutputSchema LogoutOutputSchema `json:"output_schema"`
}
