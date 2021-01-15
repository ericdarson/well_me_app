package request

type PlannerPembelianRequest struct {
	BcaID     string       `json:"bca_id" binding:"required"`
	IDPlan    int          `json:"id_plan" binding:"required"`
	KodePromo string       `json:"kode_promo"`
	Products  []ProdukBeli `json:"products" binding:"required"`
}

type ProdukBeli struct {
	IDProduk int     `json:"id_produk" binding:"required"`
	Nominal  float64 `json:"nominal" binding:"required"`
}
