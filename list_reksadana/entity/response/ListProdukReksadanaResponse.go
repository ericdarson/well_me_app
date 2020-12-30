package response

type ListProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                       `json:"error_schema"`
	OutputSchema []ListProdukReksadanaOutputSchema `json:"output_schema"`
}

type ListProdukReksadanaOutputSchema struct {
	ID      string `json:"id_jenis_reksadana"`
	Nama    string `json:"nama_jenis_reksadana"`
	Nab     string `json:"nab"`
	Kinerja string `json:"kinerja_satu_bulan"`
	DateNab string `json:"date_nab"`
}
