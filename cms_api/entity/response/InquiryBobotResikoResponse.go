package response

type InquiryBobotResikoResponse struct {
	ErrorSchema  ErrorSchema                      `json:"error_schema"`
	OutputSchema []InquiryBobotResikoOutputSchema `json:"output_schema"`
}

type InquiryBobotResikoOutputSchema struct {
	BobotResiko       string              `json:"bobot_resiko"`
	DetailBobotResiko []DetailBobotResiko `json:"detail_bobot_resiko"`
}

type DetailBobotResiko struct {
	IDJenis    string  `json:"id_jenis_reksadana"`
	NamaJenis  string  `json:"nama_jenis_reksadana"`
	Persentase float64 `json:"persentase"`
}
