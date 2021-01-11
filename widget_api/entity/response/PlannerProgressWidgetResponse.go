package response

type PlannerProgressWidgetResponse struct {
	ErrorSchema  ErrorSchema                         `json:"error_schema"`
	OutputSchema []PlannerProgressWidgetOutputSchema `json:"output_schema"`
}

type PlannerProgressWidgetOutputSchema struct {
	IdPlan     string `json:"idPlan"`
	NamaPlan   string `json:"namaPlan"`
	Percentage string `json:"percentage"`
	Target     string `json:"target"`
	Kategori   string `json:"kategori"`
}
