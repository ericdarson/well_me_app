package request

type ProdukReksadanaRequest struct {
	Nama             string  `json:"nama_produk" binding:"required"`
	IDJenisReksadana int     `json:"id_jenis_reksadana" binding:"required"`
	MinimumPembelian float64 `json:"minimum_pembelian" binding:"required"`
	ExpenseRatio     float64 `json:"expense_ratio" binding:"required"`
	TotalAUM         float64 `json:"total_aum" binding:"required"`
	ManagerInvestasi string  `json:"manager_investasi" binding:"required"`
	TingkatResiko    string  `json:"tingkat_resiko" binding:"required"`
	LevelResiko      int     `json:"level_resiko" binding:"required"`
	BankKustodian    string  `json:"bank_kustodian" binding:"required"`
	BankPenampung    string  `json:"bank_penampung" binding:"required"`
	URLVendor        string  `json:"url_vendor" binding:"required"`
	PwVendor         string  `json:"password_vendor_md5" binding:"required"`
	BiayaPembelian   float64 `json:"biaya_pembelian" binding:"required"`
	BiayaPenjualan   float64 `json:"biaya_penjualan" binding:"required"`
	MinimumSisaUnit  float64 `json:"minimum_sisa_unit" binding:"required"`
}
