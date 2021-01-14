package response

type DashboardOverviewResponse struct {
	ErrorSchema  ErrorSchema                   `json:"error_schema"`
	OutputSchema DashboardOverviewOutputSchema `json:"output_schema"`
}

type DashboardOverviewOutputSchema struct {
	User            int         `json:"user"`
	NewUser         int         `json:"new_user""`
	JumlahInvestasi float64     `json:"jumlah_investasi"`
	NewPlanner      int         `json:"new_planner"`
	ChartPembelian  []ChartData `json:"chart_pembelian"`
	ChartPenjualan  []ChartData `json:"chart_penjualan"`
}

type ChartData struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}
