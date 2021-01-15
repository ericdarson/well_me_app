package response

type InquiryDataNasabahResponse struct {
	ErrorSchema  ErrorSchema                    `json:"error_schema"`
	OutputSchema InquiryDataNasabahOutputSchema `json:"output_schema"`
}

type InquiryDataNasabahOutputSchema struct {
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	NoRekening    string `json:"no_rekening"`
	SID           string `json:"sid"`
	TglJoin       string `json:"tanggal_join"`
	WrongAttempt  string `json:"wrong_attempt"`
	DateLocked    string `json:"date_locked"`
	BobotResiko   string `json:"bobot_resiko"`
	LevelResiko   string `json:"level_resiko"`
	TingkatResiko string `json:"tingkat_resiko"`
	NoHP          string `json:"no_hp"`
}
