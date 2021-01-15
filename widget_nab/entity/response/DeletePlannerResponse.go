package response

type DeletePlannerResponse struct {
	ErrorSchema  ErrorSchema               `json:"error_schema"`
	OutputSchema DeletePlannerOutputSchema `json:"output_schema"`
}
