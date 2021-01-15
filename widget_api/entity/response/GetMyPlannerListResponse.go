package response

type GetMyPlannerListResponse struct {
	ErrorSchema  ErrorSchema                 `json:"error_schema"`
	OutputSchema []MyPlannerListOutputSchema `json:"output_schema"`
}

type MyPlannerListOutputSchema struct {
	IdPlan   string `json:"idPlan"`
	NamaPlan string `json:"namaPlan"`
}
