package request

type UpdateBobotResikoRequest struct {
	Input []UpdateBobotResikoSingleRequest `json:"input" binding:"required"`
}

type UpdateBobotResikoSingleRequest struct {
	BobotResiko float64 `json:"bobot_resiko" binding:"required"`
	IDJenis     int     `json:"id_jenis_reksadana" binding:"required"`
	Persentase  float64 `json:"persentase" binding:"required"`
}
