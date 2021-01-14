package response

type TransactionHistoryResponse struct {
	ErrorSchema  ErrorSchema                      `json:"error_schema"`
	OutputSchema []TransactionHistoryOutputSchema `json:"output_schema"`
}

type TransactionHistoryOutputSchema struct {
	TransactionID              string  `json:"transaction_id"`
	BCAID                      string  `json:"bca_id"`
	IDProduk                   string  `json:"id_produk"`
	NamaProduk                 string  `json:"nama_produk"`
	IDPlan                     string  `json:"id_plan"`
	NamaPlan                   string  `json:"nama_plan"`
	Status                     string  `json:"status_transaksi"`
	KodePromo                  string  `json:"kode_promo"`
	Nab                        float64 `json:"nab"`
	JumlahUnit                 float64 `json:"jumlah_unit"`
	TotalNominal               float64 `json:"total_nominal"`
	TanggalTransaksi           string  `json:"tanggal_transaksi"`
	TanggalVerifikasiBank      string  `json:"tanggal_verifikasi_bank"`
	TanggalVerifikasiPembelian string  `json:"tanggal_verifikasi_pembelian"`
}
