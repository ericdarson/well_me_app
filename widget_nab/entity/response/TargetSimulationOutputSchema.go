package response

type TargetSimulationResponse struct {
	ErrorSchema  ErrorSchema                  `json:"error_schema"`
	OutputSchema TargetSimulationOutputSchema `json:"output_schema"`
}
