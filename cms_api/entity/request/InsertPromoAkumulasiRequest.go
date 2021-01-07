package request

type InsertPromoAkumulasiRequest struct {
	Title       string  `json:"title" binding:"required"`
	Subtitle    string  `json:"subtitle" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     string  `json:"end_date" binding:"required"`
	Description string  `json:"description"`
	CashBack    float64 `json:"cashback" binding:"required"`
	Target      float64 `json:"target_akumulasi" binding:"required"`
}
