package response

type InquiryPromoResponse struct {
	ErrorSchema  ErrorSchema                `json:"error_schema"`
	OutputSchema []InquiryPromoOutputSchema `json:"output_schema"`
}

type InquiryPromoOutputSchema struct {
	KodePromo    string `json:"kode_promo"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Description  string `json:"description"`
	Target       string `json:"target_akumulasi"`
	Cashback     string `json:"cashback"`
	Minimum      string `json:"minimum_transaction"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	PromoType    string `json:"promo_type"`
	ActiveStatus string `json:"active_status"`
}
