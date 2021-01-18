package response

type UpdatePlannerResponse struct {
	ErrorSchema  ErrorSchema               `json:"error_schema"`
	OutputSchema UpdatePlannerOutputSchema `json:"output_schema"`
}
