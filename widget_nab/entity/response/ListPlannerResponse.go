package response

type ListPlannerResponse struct {
	ErrorSchema  ErrorSchema             `json:"error_schema"`
	OutputSchema ListPlannerOutputSchema `json:"output_schema"`
}
