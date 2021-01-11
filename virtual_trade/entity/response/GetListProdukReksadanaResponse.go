package response

type GetListProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                          `json:"error_schema"`
	OutputSchema []GetListProdukReksadanaOutputSchema `json:"output_schema"`
}

type GetListProdukReksadanaOutputSchema struct {
	Id               string `json:"id_produk"`
	Nama             string `json:"nama"`
	KinerjaSatuBulan string `json:"kinerja_satu_bulan"`
	Nab              string `json:"nab"`
	MaxBackwardDate  string `json:"max_backward_date"`
}
