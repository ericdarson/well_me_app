package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type NabWidgetDao interface {
	GetByIds(string) []response.DetailNAB
}

type nabWidgetDao struct {
	temp string
}

func New() NabWidgetDao {
	return &nabWidgetDao{}
}

func (dao *nabWidgetDao) GetByIds(listIds string) []response.DetailNAB {
	var listNAB []response.DetailNAB
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.DetailNAB{
			Reksadana: "-1",
			Nab:       fmt.Sprint(err),
		}
		listNAB = append(listNAB, single)
		return listNAB
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetNabByIds.query")
	if err != nil {
		fmt.Println(err)
		single := response.DetailNAB{
			Reksadana: "-1",
			Nab:       fmt.Sprint(err),
		}
		listNAB = append(listNAB, single)
		return listNAB
	}
	query := string(dat)
	query = fmt.Sprintf(query, listIds)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error running query")
		single := response.DetailNAB{
			Reksadana: "-1",
			Nab:       fmt.Sprint(err),
		}
		listNAB = append(listNAB, single)
		return listNAB
	}
	defer rows.Close()

	for rows.Next() {

		var reksadana string
		var nab string
		rows.Scan(&reksadana, &nab)
		single := response.DetailNAB{
			Reksadana: reksadana,
			Nab:       nab,
		}

		listNAB = append(listNAB, single)
	}

	return listNAB
}

/*
func main() {

	db, err := sql.Open("godror", `user="test" password="test123" connectString="localhost:1521/xe"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
*/
