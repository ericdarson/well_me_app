package response

type GetJenisReksadanaResponse struct {
	ErrorSchema  ErrorSchema                     `json:"error_schema"`
	OutputSchema []GetJenisReksadanaOutputSchema `json:"output_schema"`
}

type GetJenisReksadanaOutputSchema struct {
	Id        string `json:"id_jenis_reksadana"`
	Reksadana string `json:"jenis_reksadana"`
}
