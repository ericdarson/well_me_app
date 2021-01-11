package response

type GetDetailProdukReksadanaResponse struct {
	ErrorSchema  ErrorSchema                          `json:"error_schema"`
	OutputSchema GetDetailProdukReksadanaOutputSchema `json:"output_schema"`
}

type GetDetailProdukReksadanaOutputSchema struct {
	ID               string     `json:"id_produk"`
	Nama             string     `json:"nama_produk"`
	CagrOneWeek      float64    `json:"cagr_one_week"`
	CagrOneMonth     float64    `json:"cagr_one_month"`
	CagrThreeMonths  float64    `json:"cagr_three_months"`
	CagrOneYear      float64    `json:"cagr_one_year"`
	Nab              float64    `json:"nab"`
	Expratio         string     `json:"expense_ratio"`
	Aum              string     `json:"aum"`
	ManagerInvestasi string     `json:"manager_investasi"`
	Resiko           string     `json:"resiko"`
	Minimal          string     `json:"minimal_pembelian"`
	BankKustodian    string     `json:"bank_kustodian"`
	BankPenampung    string     `json:"bank_penampung"`
	SystemDate       string     `json:"system_date"`
	SystemDateString string     `json:"system_date_string"`
	NabOneWeek       []DailyNab `json:"one_week_daily_nab"`
	NabOneMonth      []DailyNab `json:"one_month_daily_nab"`
	NabThreeMonths   []DailyNab `json:"three_months_daily_nab"`
	NabOneYear       []DailyNab `json:"one_year_daily_nab"`
}

type DailyNab struct {
	Date       string  `json:"date"`
	DateString string  `json:"datestring"`
	Nab        float64 `json:"nab_daily"`
}
