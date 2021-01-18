package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
)

type UpdatePlannerDao interface {
	UpdatePlanner(ctx *gin.Context, idPlan string, namaPlan string, periodic string, due_date string, kategori string) response.DetailUpdatePlanner
}

type updatePlannerDao struct {
	temp string
}

func NewUpdatePlannerDao() UpdatePlannerDao {
	return &updatePlannerDao{}
}
func (dao *updatePlannerDao) UpdatePlanner(ctx *gin.Context, idPlan string, namaPlan string, periodic string, due_date string, kategori string) response.DetailUpdatePlanner {
	var detailUpdatePlanner response.DetailUpdatePlanner
	conn := dbconnection.New()
	db := conn.GetConnection()

	var message string

	query := `BEGIN SP_UPDATE_COLUMN_PLANNER ('` + idPlan + `','` + namaPlan + `','` + periodic + `','` + due_date + `','` + kategori + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &message}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return detailUpdatePlanner
	} else {
		if message == "1" {
			detailUpdatePlanner.Message = "SUKSES"

		} else {
			detailUpdatePlanner.Message = message
		}
	}
	return detailUpdatePlanner

}
