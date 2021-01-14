package response

type JenisReksadanaResponse struct {
	ErrorSchema  ErrorSchema                  `json:"error_schema"`
	OutputSchema []JenisReksadanaOutputSchema `json:"output_schema"`
}

type JenisReksadanaOutputSchema struct {
	ID   string `json:"id_jenis"`
	Nama string `json:"nama_jenis"`
}
