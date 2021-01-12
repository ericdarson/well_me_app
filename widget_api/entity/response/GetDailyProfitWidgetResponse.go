package response

type GetDailyProfitWidgetResponse struct {
	ErrorSchema  ErrorSchema                        `json:"error_schema"`
	OutputSchema []GetDailyProfitWidgetOutputSchema `json:"output_schema"`
}

type GetDailyProfitWidgetOutputSchema struct {
	Id               string      `json:"id_produk"`
	Nama             string      `json:"nama"`
	CurrBalance      string      `json:"current_balance"`
	CurrProfit       string      `json:"current_profit"`
	ChartDataOneYear []ChartData `json:"chart_data_one_year"`
}

type ChartData struct {
	DateDaily       string `json:"date_daily"`
	DateDailyString string `json:"datestring_daily"`
	NabDaily        string `json:"nab_daily"`
}
