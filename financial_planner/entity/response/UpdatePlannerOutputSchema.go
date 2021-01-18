package response

type UpdatePlannerOutputSchema struct {
	SystemDate          string              `json:"system_date"`
	DetailUpdatePlanner DetailUpdatePlanner `json:"detail_update_planner"`
}

type DetailUpdatePlanner struct {
	Message string `json:"message"`
}
