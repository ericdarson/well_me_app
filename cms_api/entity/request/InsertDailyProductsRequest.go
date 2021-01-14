package request

type InsertDailyProductsRequest struct {
	Input []InsertDailyProduct `json:"input" binding:"required"`
}

type InsertDailyProduct struct {
	IDProduk int     `json:"id_produk" binding:"required"`
	Nab      float64 `json:"nab" binding:"required"`
}
