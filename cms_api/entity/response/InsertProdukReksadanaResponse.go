package response

type InsertProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                       `json:"error_schema"`
	OutputSchema InsertProdukReksadanaOutputSchema `json:"output_schema"`
}

type InsertProdukReksadanaOutputSchema struct {
	ID string `json:"id"`
}
