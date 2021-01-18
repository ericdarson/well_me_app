package response

type TargetSimulationOutputSchema struct {
	SystemDate             string                   `json:"system_date"`
	TargetSimulationDetail []TargetSimulationDetail `json:"detail_chart"`
	TabunganPeriode        string                   `json:"nominal_per_periode"`
}

type TargetSimulationDetail struct {
	NominalTanpaInvestasi string
	NominalInvestasi      string
	Date                  string
}
