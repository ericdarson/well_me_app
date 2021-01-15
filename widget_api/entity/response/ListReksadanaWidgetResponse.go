package response

type ListReksadanaWidgetResponse struct {
	ErrorSchema  ErrorSchema                       `json:"error_schema"`
	OutputSchema []ListReksadanaWidgetOutputSchema `json:"output_schema"`
}

type ListReksadanaWidgetOutputSchema struct {
	Id     string `json:"id"`
	Nama   string `json:"reksadana"`
	Nab    string `json:"nab"`
	Profit string `json:"profit"`
}
