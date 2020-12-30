package dao

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type ListPlannerDao interface {
	GetById(string) []response.DetailPlanner
}

type listPlannerDao struct {
	temp string
}

func NewListPlannerDao() ListPlannerDao {
	return &listPlannerDao{}
}

func (dao *listPlannerDao) GetById(BcaId string) []response.DetailPlanner {
	var listPlanner []response.DetailPlanner
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.DetailPlanner{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	dat, err := ioutil.ReadFile(dir + "\\query\\getListPlanner.query")
	if err != nil {
		fmt.Println(err)
		single := response.DetailPlanner{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	query := string(dat)
	query = fmt.Sprintf(query, BcaId)

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error running %q: %+v", query, err)
		single := response.DetailPlanner{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	defer rows.Close()

	for rows.Next() {
		var IdPlan string
		var NamaPlan string
		var Percentage string
		var Target string
		var Kategori string

		rows.Scan(&IdPlan, &NamaPlan, &Percentage, &Target, &Kategori)
		var message string
		query := `BEGIN SP_UPDATE_PLANNER('` + IdPlan + `','` + BcaId + `',:1); END;`
		if _, err := db.Exec(query, sql.Out{Dest: &message}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
		}
		Percentage = Percentage + "%"
		single := response.DetailPlanner{
			IdPlan:     IdPlan,
			NamaPlan:   NamaPlan,
			Percentage: Percentage,
			Target:     Target,
			Kategori:   Kategori,
		}

		listPlanner = append(listPlanner, single)
	}

	return listPlanner
}
