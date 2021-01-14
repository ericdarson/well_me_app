package response

type InquiryBobotResikoResponse struct {
	ErrorSchema  ErrorSchema                    `json:"error_schema"`
	OutputSchema map[string][]DetailBobotResiko `json:"output_schema"`
}

type DetailBobotResiko struct {
	IDJenis    string  `json:"id_jenis_reksadana"`
	NamaJenis  string  `json:"nama_jenis_reksadana"`
	Persentase float64 `json:"persentase"`
}
