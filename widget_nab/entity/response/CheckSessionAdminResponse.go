package response

type CheckSessionAdminResponse struct {
	ErrorSchema  ErrorSchema                   `json:"error_schema"`
	OutputSchema CheckSessionAdminOutputSchema `json:"output_schema"`
}
