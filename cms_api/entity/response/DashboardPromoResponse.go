package response

type DashboardPromoResponse struct {
	ErrorSchema  ErrorSchema                `json:"error_schema"`
	OutputSchema DashboardPromoOutputSchema `json:"output_schema"`
}

type DashboardPromoOutputSchema struct {
	Objectives []Objective `json:"objectives"`
	Promotions []Promotion `json:"promotions"`
}

type Promotion struct {
	KodePromo string `json:"kode_promo"`
	Title     string `json:"title"`
	Used      int    `json:"use_qty"`
}

type Objective struct {
	KodePromo string `json:"kode_promo"`
	Title     string `json:"title"`
	Started   int    `json:"start_qty"`
	Claimed   int    `json:"claim_qty"`
}
