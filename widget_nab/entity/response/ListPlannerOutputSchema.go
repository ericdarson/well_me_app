package response

type ListPlannerOutputSchema struct {
	SystemDate  string          `json:"system_date"`
	ListPlanner []DetailPlanner `json:"list_planner"`
}

type DetailPlanner struct {
	IdPlan     string `json:"idPlan"`
	NamaPlan   string `json:"namaPlan"`
	Percentage string `json:"percentage"`
	Target     string `json:"target"`
	Kategori   string `json: "kategori"`
}
