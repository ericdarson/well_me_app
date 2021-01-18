package response

type DeletePlannerOutputSchema struct {
	SystemDate          string              `json:"system_date"`
	DetailDeletePlanner DetailDeletePlanner `json:"detail_delete_planner"`
}

type DetailDeletePlanner struct {
	Message string `json:"message"`
}
