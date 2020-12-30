package response

type SimulationStartResponse struct {
	ErrorSchema  ErrorSchema                 `json:"error_schema"`
	OutputSchema SimulationStartOutputSchema `json:"output_schema"`
}

type SimulationStartOutputSchema struct {
	JumlahUnit  float64 `json:"jumlah_unit"`
	StartingNab string  `json:"starting_nab"`
	StartDate   string  `json:"start_date"`
}
