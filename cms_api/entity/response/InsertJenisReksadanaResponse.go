package response

type InsertJenisReksadanaResponse struct {
	ErrorSchema  ErrorSchema                `json:"error_schema"`
	OutputSchema JenisReksadanaOutputSchema `json:"output_schema"`
}
