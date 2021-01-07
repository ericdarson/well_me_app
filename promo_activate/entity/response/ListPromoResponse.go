package response

type ListPromoResposne struct {
	ErrorSchema  ErrorSchema           `json:"error_schema"`
	OutputSchema ListPromoOutputSchema `json:"output_schema"`
}

type ListPromoOutputSchema struct {
	Objectives []Objectives `json:"objectives"`
	Promotion  []Promotion  `json:"promotions"`
}

type Objectives struct {
	KodePromo     string `json:"kode_promo"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Description   string `json:"description"`
	Current       string `json:"current_amount"`
	Target        string `json:"target_akumulasi"`
	Cashback      string `json:"cashback"`
	DateAvailable string `json:"date_available"`
}

type Promotion struct {
	KodePromo     string `json:"kode_promo"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Description   string `json:"description"`
	Minimum       string `json:"minimum_transaction"`
	Cashback      string `json:"cashback"`
	DateAvailable string `json:"date_available"`
}
