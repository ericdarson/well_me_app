package request

type InsertPromoTransaksiRequest struct {
	KodePromo   string  `json:"kode_promo" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Subtitle    string  `json:"subtitle" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     string  `json:"end_date" binding:"required"`
	Description string  `json:"description" binding:"required"`
	CashBack    float64 `json:"cashback" binding:"required"`
	Minimum     float64 `json:"minimum_transaction" binding:"required"`
}
