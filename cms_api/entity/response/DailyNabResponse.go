package response

type DailyNabResponse struct {
	ErrorSchema  ErrorSchema            `json:"error_schema"`
	OutputSchema []DailyNabOutputSchema `json:"output_schema"`
}
type DailyNabOutputSchema struct {
	IDProduk       int     `json:"id_produk"`
	NamaProduk     string  `json:"nama_produk"`
	Nab            float64 `json:"latest_nab"`
	IsUpdatedToday int     `json:"is_updated_today"`
}
