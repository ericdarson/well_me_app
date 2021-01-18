package response

type LoginResponse struct {
	ErrorSchema  ErrorSchema       `json:"error_schema"`
	OutputSchema LoginOutputSchema `json:"output_schema"`
}
