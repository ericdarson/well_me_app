package dao

import (
	"database/sql"
	"fmt"
	"log"
	"widget_nab_service/dbconnection"

	_ "github.com/godror/godror"
)

type InsertPlannerDao interface {
	InsertPlanner(bcaId string, namaPlan string, goalAmount string, currentAmount string, periodic string, dueDate string, kategori string) string
}

type insertPlannerDao struct {
	temp string
}

func NewInsertPlannerDao() InsertPlannerDao {
	return &insertPlannerDao{}
}
func (dao *insertPlannerDao) InsertPlanner(bcaId string, namaPlan string, goalAmount string, currentAmount string, periodic string, dueDate string, kategori string) string {
	var detailInsertPlanner string
	conn := dbconnection.New()
	db := conn.GetConnection()

	var status string
	fmt.Println(goalAmount)

	query := `BEGIN SP_INSERT_PLANNER('` + bcaId + `','` + namaPlan + `','` + goalAmount + `','` + currentAmount + `','` + periodic + `','` + dueDate + `','` + bcaId + `','` + kategori + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &status}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return detailInsertPlanner
	} else {
		if status == "1" {
			detailInsertPlanner = "Sukses"
		} else {
			detailInsertPlanner = "Gagal"
		}
	}
	return detailInsertPlanner

}
