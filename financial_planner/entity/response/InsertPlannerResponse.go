package response

type InsertPlannerResponse struct {
	ErrorSchema  ErrorSchema               `json:"error_schema"`
	OutputSchema InsertPlannerOutputSchema `json:"output_schema"`
}
