package handler

import (
	"fmt"
	"strconv"
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	updatePlannerDaoObject dao.UpdatePlannerDao = dao.NewUpdatePlannerDao()
)

//idPlan string, namaPlan string, periodic string, due_date string
type UpdatePlannerRequest struct {
	IDPlan   int    `json:"idPlan"`
	NamaPlan string `json:"namaPlan"`
	Periodic string `json:"periodic"`
	DueDate  string `json:"dueDate"`
	Kategori string `json:"kategori"`
}

func UpdatePlanner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		updatePlannerRequest := UpdatePlannerRequest{}
		ctx.Bind(&updatePlannerRequest)
		fmt.Println(updatePlannerRequest.IDPlan)
		currentTime := time.Now()
		var updatePlannerResponse response.UpdatePlannerResponse
		var outputSchema response.UpdatePlannerOutputSchema
		var errorSchema response.ErrorSchema
		var idPlan string
		idPlan = strconv.Itoa(updatePlannerRequest.IDPlan)
		if updatePlannerRequest.IDPlan == 0 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := updatePlannerDaoObject.UpdatePlanner(ctx, idPlan, updatePlannerRequest.NamaPlan, updatePlannerRequest.Periodic, updatePlannerRequest.DueDate, updatePlannerRequest.Kategori)
			if result.Message == "" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailUpdatePlanner = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				updatePlannerResponse.OutputSchema = outputSchema
				responseCode = 200
			}
		}
		updatePlannerResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, updatePlannerResponse)
	}
}
