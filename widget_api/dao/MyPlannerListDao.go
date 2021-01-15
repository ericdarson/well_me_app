package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"widget_api/dbconnection"
	"widget_api/entity/response"

	_ "github.com/godror/godror"
)

type MyPlannerListDao interface {
	GetMyPlannerList(string, string, string) []response.MyPlannerListOutputSchema
}

type myPlannerListDao struct {
	temp string
}

func NewMyPlannerList() MyPlannerListDao {
	return &myPlannerListDao{}
}

func (dao *myPlannerListDao) GetMyPlannerList(bcaid string, name string, sortBy string) []response.MyPlannerListOutputSchema {
	var listPlanner []response.MyPlannerListOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.MyPlannerListOutputSchema{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	dat, err := ioutil.ReadFile(dir + "/query/getMyPlannerList.query")
	if err != nil {
		fmt.Println(err)
		single := response.MyPlannerListOutputSchema{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	var filterBy string
	var orderBy string
	tempSort, err := strconv.Atoi(sortBy)
	if err != nil {
		fmt.Println(err)
		filterBy = "t.nama_produk"
		orderBy = "asc"
	} else {
		if tempSort == 1 {
			filterBy = "NAMA_PLAN"
			orderBy = "asc"
		} else if tempSort == 2 {
			filterBy = "NAMA_PLAN"
			orderBy = "desc"
		}
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaid, name, filterBy, orderBy)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.MyPlannerListOutputSchema{
			IdPlan: "-1",
		}
		listPlanner = append(listPlanner, single)
		return listPlanner
	}
	defer rows.Close()

	for rows.Next() {
		var IdPlan string
		var NamaPlan string

		rows.Scan(&IdPlan, &NamaPlan)
		single := response.MyPlannerListOutputSchema{
			IdPlan:   IdPlan,
			NamaPlan: NamaPlan,
		}
		listPlanner = append(listPlanner, single)
	}

	return listPlanner
}
