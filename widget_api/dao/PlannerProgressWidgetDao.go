package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_api/dbconnection"
	"widget_api/entity/response"

	_ "github.com/godror/godror"
)

type ProgressPlannerWidgetDao interface {
	GetProgressPlannerWidget(string) []response.PlannerProgressWidgetOutputSchema
}

type progressPlannerWidgetDao struct {
	temp string
}

func NewProgressPlanner() ProgressPlannerWidgetDao {
	return &progressPlannerWidgetDao{}
}

func (dao *progressPlannerWidgetDao) GetProgressPlannerWidget(listIds string) []response.PlannerProgressWidgetOutputSchema {
	var listPlanner []response.PlannerProgressWidgetOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.PlannerProgressWidgetOutputSchema{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	dat, err := ioutil.ReadFile(dir + "/query/getPlannerProgress.query")
	if err != nil {
		fmt.Println(err)
		single := response.PlannerProgressWidgetOutputSchema{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	query := string(dat)
	query = fmt.Sprintf(query, listIds)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.PlannerProgressWidgetOutputSchema{
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
		Percentage = Percentage + "%"
		single := response.PlannerProgressWidgetOutputSchema{
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
