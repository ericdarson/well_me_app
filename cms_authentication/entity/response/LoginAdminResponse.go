package response

type LoginAdminResponse struct {
	ErrorSchema  ErrorSchema            `json:"error_schema"`
	OutputSchema LoginAdminOutputSchema `json:"output_schema"`
}
