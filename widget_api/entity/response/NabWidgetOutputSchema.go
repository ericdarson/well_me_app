package response

type NabWidgetOutputSchema struct {
	SystemDate string      `json:"system_date"`
	ListNAB    []DetailNAB `json:"list_nab"`
}

type DetailNAB struct {
	Reksadana string `json:"reksadana"`
	Nab       string `json:"nab"`
}
