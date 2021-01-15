package handler

import (
	"log"
	"strconv"
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	insertPlannerDaoObject dao.InsertPlannerDao = dao.NewInsertPlannerDao()
)

//bcaId string, namaPlan string, goalAmount string, currentAmount string, periodic string, dueDate string
type InsertPlannerRequest struct {
	BcaId      string `json:"bcaId"`
	NamaPlan   string `json:"namaPlan"`
	GoalAmount string `json:"goalAmount"`
	Periodic   string `json:"periodic"`
	DueDate    string `json:"dueDate"`
	Kategori   string `json:"kategori"`
}

func InsertPlanner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		insertPlannerRequest := InsertPlannerRequest{}
		err := ctx.Bind(&insertPlannerRequest)
		_, err2 := strconv.ParseFloat(insertPlannerRequest.GoalAmount, 32)
		layout := "02-01-2006"
		_, err3 := time.Parse(layout, insertPlannerRequest.DueDate)
		currentTime := time.Now()

		var insertPlannerResponse response.InsertPlannerResponse
		var outputSchema response.InsertPlannerOutputSchema
		var errorSchema response.ErrorSchema
		if err2 != nil || err != nil || err3 != nil || insertPlannerRequest.BcaId == "" || insertPlannerRequest.NamaPlan == "" || (insertPlannerRequest.Periodic != "Weekly" && insertPlannerRequest.Periodic != "Monthly" && insertPlannerRequest.Periodic != "Yearly") || insertPlannerRequest.Kategori == "" {
			log.Printf("Error : %+v", err3, err, err2, insertPlannerRequest.BcaId, insertPlannerRequest.NamaPlan, insertPlannerRequest.Periodic)
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := insertPlannerDaoObject.InsertPlanner(insertPlannerRequest.BcaId, insertPlannerRequest.NamaPlan, insertPlannerRequest.GoalAmount, "0", insertPlannerRequest.Periodic, insertPlannerRequest.DueDate, insertPlannerRequest.Kategori)
			if result == "Gagal" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.Message = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				insertPlannerResponse.OutputSchema = outputSchema
				responseCode = 200
			}
		}
		insertPlannerResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, insertPlannerResponse)
	}
}
