package response

type PlannerGamificationResponse struct {
	ErrorSchema  ErrorSchema                     `json:"error_schema"`
	OutputSchema PlannerGamificationOutputSchema `json:"output_schema"`
}

type PlannerGamificationOutputSchema struct {
	Nama    string `json:"nama_plan"`
	Target  int    `json:"target_plan"`
	Amount  int    `json:"current_amount"`
	Gambar  string `json:"gambar"`
	Puzzle  string `json:"puzzle_sequence"`
	DueDate string `json:"due_date"`
}
