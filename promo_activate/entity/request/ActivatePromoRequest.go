package request

type ActivatePromoRequest struct {
	BcaID     string `json:"bca_id" binding:"required"`
	KodePromo string `json:"kode_promo" binding:"required"`
}
