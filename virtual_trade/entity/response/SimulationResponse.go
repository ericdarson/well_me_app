package response

type SimulationResponse struct {
	ErrorSchema  ErrorSchema            `json:"error_schema"`
	OutputSchema SimulationOutputSchema `json:"output_schema"`
}

type SimulationOutputSchema struct {
	DateSimulation       string `json:"nab_date"`
	DateSimulationString string `json:"nab_datestring"`
	NabSimulation        string `json:"nab_simulation"`
}
