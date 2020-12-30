package response

type ProfileResponse struct {
	ErrorSchema  ErrorSchema         `json:"error_schema"`
	OutputSchema ProfileOutputSchema `json:"output_schema"`
}
