package response

type NabWidgetResponse struct {
	ErrorSchema  ErrorSchema           `json:"error_schema"`
	OutputSchema NabWidgetOutputSchema `json:"output_schema"`
}
