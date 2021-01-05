//Nama, Profile Resiko, BCA ID, Email, SID, No Rekening BCA

package response

type ProfileOutputSchema struct {
	SystemDate    string        `json:"system_date"`
	DetailProfile DetailProfile `json:"detail_profile"`
}

type DetailProfile struct {
	BcaId         string `json:"bcaId"`
	Nama          string `json:"nama"`
	ProfileResiko string `json:"profileResiko"`
	Email         string `json:"email"`
	SID           string `json:"sid"`
	NoRekBCA      string `json:"noRekBCA"`
}
