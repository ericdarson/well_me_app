package request

type InsertJenisReksadanaRequest struct {
	Nama string `json:"nama_jenis_reksadana" binding:"required"`
}
