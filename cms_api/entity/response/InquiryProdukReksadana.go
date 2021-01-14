package response

type InquiryProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                   `json:"error_schema"`
	OutputSchema []ProdukReksadanaOutputSchema `json:"output_schema"`
}

type ProdukReksadanaOutputSchema struct {
	ID                string  `json:"id_reksadana"`
	Nama              string  `json:"nama_reksadana"`
	Nab               float64 `json:"nab"`
	Minimum           string  `json:"minimum_pembelian"`
	ExpenseRatio      float64 `json:"expense_ratio"`
	TotalAUM          float64 `json:"total_aum"`
	ManagerInvestasi  string  `json:"manager_investasi"`
	Resiko            string  `json:"resiko"`
	LevelResiko       int     `json:"level_resiko"`
	BankKustodian     string  `json:"bank_kustodian"`
	BankPenampung     string  `json:"bank_penampung"`
	KinerjaSatuMinggu string  `json:"kinerja_satu_minggu"`
	KinerjaSatuBulan  string  `json:"kinerja_satu_bulan"`
	KinerjaTigaBulan  string  `json:"kinerja_tiga_bulan"`
	KinerjaSatuTahun  string  `json:"kinerja_satu_tahun"`
	IDJenis           string  `json:"id_jenis_reksadana"`
	NamaJenis         string  `json:"nama_jenis_reksadana"`
	URLVendor         string  `json:"url_vendor"`
	PwVendor          string  `json:"password_vendor_md5"`
}
