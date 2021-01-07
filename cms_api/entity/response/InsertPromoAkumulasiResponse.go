package response

type InsertPromoAkumulasiResponse struct {
	ErrorSchema  ErrorSchema                      `json:"error_schema"`
	OutputSchema InsertPromoAkumulasiOutputSchema `json:"output_schema"`
}

type InsertPromoAkumulasiOutputSchema struct {
	KodePromo   string  `json:"kode_promo"`
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Description string  `json:"description"`
	CashBack    float64 `json:"cashback"`
	Target      float64 `json:"target_akumulasi"`
}
