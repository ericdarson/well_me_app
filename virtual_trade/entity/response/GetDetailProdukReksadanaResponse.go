package response

type GetDetailProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                          `json:"error_schema"`
	OutputSchema GetDetailProdukReksadanaOutputSchema `json:"output_schema"`
}

type GetDetailProdukReksadanaOutputSchema struct {
	ID               string     `json:"id_produk"`
	Nama             string     `json:"nama_produk"`
	Cagr             string     `json:"cagr"`
	Nab              string     `json:"nab"`
	Expratio         string     `json:"expense_ratio"`
	Aum              string     `json:"aum"`
	ManagerInvestasi string     `json:"manager_investasi"`
	Resiko           string     `json:"resiko"`
	Minimal          string     `json:"minimal_pembelian"`
	BankKustodian    string     `json:"bank_kustodian"`
	BankPenampung    string     `json:"bank_penampung"`
	SystemDate       string     `json:"system_date"`
	NabDaily         []DailyNab `json:"daily_nab"`
}

type DailyNab struct {
	Date string `json:"date"`
	Nab  string `json:"nab_daily"`
}
