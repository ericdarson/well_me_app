package response

type PlannerResikoResponse struct {
	ErrorSchema  ErrorSchema                 `json:"error_schema"`
	OutputSchema []PlannerResikoOutputSchema `json:"output_schema"`
}

type PlannerResikoOutputSchema struct {
	IDJenis    string `json:"id_jenis_reksadana"`
	Nama       string `json:"nama_plan"`
	Percentage string `json:"percentage"`
}
