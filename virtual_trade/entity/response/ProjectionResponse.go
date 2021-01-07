package response

type ProjectionResponse struct {
	ErrorSchema  ErrorSchema            `json:"error_schema"`
	OutputSchema ProjectionOutputSchema `json:"output_schema"`
}

type ProjectionOutputSchema struct {
	ID                  string               `json:"id_produk"`
	Nama                string               `json:"nama_produk"`
	Date                string               `json:"date"`
	DateString          string               `json:"datestring"`
	Nab                 string               `json:"current_nab"`
	CagrOneYear         float64              `json:"cagr_one_year"`
	CagrThreeMonths     float64              `json:"cagr_three_months"`
	CagrOneMonth        float64              `json:"cagr_one_month"`
	CagrOneWeek         float64              `json:"cagr_one_week"`
	CartDataOneYear     []ProjectionCartData `json:"cart_data_one_year"`
	CartDataThreeMonths []ProjectionCartData `json:"cart_data_three_months"`
	CartDataOneMonth    []ProjectionCartData `json:"cart_data_one_month"`
	CartDataOneWeek     []ProjectionCartData `json:"cart_data_one_week"`
}

type ProjectionCartData struct {
	DateDaily       string `json:"date_daily"`
	DateDailyString string `json:"datestring_daily"`
	NabDaily        string `json:"nab_daily"`
}
