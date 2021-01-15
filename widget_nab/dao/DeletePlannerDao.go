package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type DeletePlannerDao interface {
	DeletePlanner(idPlan string, bcaId string) response.DetailDeletePlanner
}

type deletePlannerDao struct {
}

func NewDeletePlannerDao() DeletePlannerDao {
	return &deletePlannerDao{}
}

func (dao *deletePlannerDao) DeletePlanner(idPlan string, bcaId string) response.DetailDeletePlanner {
	conn := dbconnection.New()
	db := conn.GetConnection()
	entityResponse := response.DetailDeletePlanner{}
	query := `BEGIN SP_DELETE_PLANNER('` + idPlan + `','` + bcaId + `',:1); END;`

	if _, err := db.Exec(query, sql.Out{Dest: &entityResponse.Message}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return entityResponse
	} else {
		entityResponse.Message = "SUKSES"
		return entityResponse
	}
}
