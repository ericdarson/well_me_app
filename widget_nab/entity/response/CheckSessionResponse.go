package response

type CheckSessionResponse struct {
	ErrorSchema  ErrorSchema              `json:"error_schema"`
	OutputSchema CheckSessionOutputSchema `json:"output_schema"`
}
