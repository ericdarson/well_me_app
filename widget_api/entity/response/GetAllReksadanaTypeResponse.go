package response

type GetAllReksadanaTypeResponse struct {
	ErrorSchema  ErrorSchema                       `json:"error_schema"`
	OutputSchema []GetAllReksadanaTypeOutputSchema `json:"output_schema"`
}

type GetAllReksadanaTypeOutputSchema struct {
	Id            string `json:"id"`
	NamaReksadana string `json:"namaReksadana"`
}
