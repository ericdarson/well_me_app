package response

type SellTransactionHistoryResponse struct {
	ErrorSchema  ErrorSchema                          `json:"error_schema"`
	OutputSchema []SellTransactionHistoryOutputSchema `json:"output_schema"`
}

type SellTransactionHistoryOutputSchema struct {
	TransactionID              string  `json:"transaction_id"`
	BCAID                      string  `json:"bca_id"`
	IDProduk                   string  `json:"id_produk"`
	NamaProduk                 string  `json:"nama_produk"`
	IDPlan                     string  `json:"id_plan"`
	NamaPlan                   string  `json:"nama_plan"`
	Status                     string  `json:"status_transaksi"`
	Nab                        float64 `json:"nab"`
	JumlahUnit                 float64 `json:"jumlah_unit"`
	TotalNominal               float64 `json:"total_nominal"`
	TanggalTransaksi           string  `json:"tanggal_transaksi"`
	TanggalVerifikasiBank      string  `json:"tanggal_verifikasi_bank"`
	TanggalVerifikasiPembelian string  `json:"tanggal_verifikasi_pembelian"`
}
