package response

import "cms_api/entity/request"

type InsertDailyProductsResponse struct {
	ErrorSchema  ErrorSchema                     `json:"error_schema"`
	OutputSchema InsertDailyProductsOutputSchema `json:"output_schema"`
}

type InsertDailyProductsOutputSchema struct {
	ListGagal []request.InsertDailyProduct `json:"list_gagal"`
}
