package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_api/dbconnection"
	"widget_api/entity/response"

	_ "github.com/godror/godror"
)

type AllReksadanaTypeDao interface {
	GetAllType() []response.GetAllReksadanaTypeOutputSchema
}

type allReksadanaTypeDao struct {
	temp string
}

func NewAllReksadanaType() AllReksadanaTypeDao {
	return &allReksadanaTypeDao{}
}

func (dao *allReksadanaTypeDao) GetAllType() []response.GetAllReksadanaTypeOutputSchema {
	var listType []response.GetAllReksadanaTypeOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.GetAllReksadanaTypeOutputSchema{
			Id:            "-1",
			NamaReksadana: fmt.Sprint(err),
		}
		listType = append(listType, single)
		return listType
	}
	dat, err := ioutil.ReadFile(dir + "/query/getAllReksadanaType.query")
	if err != nil {
		fmt.Println(err)
		single := response.GetAllReksadanaTypeOutputSchema{
			Id:            "-1",
			NamaReksadana: fmt.Sprint(err),
		}
		listType = append(listType, single)
		return listType
	}
	query := string(dat)
	query = fmt.Sprintf(query)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		single := response.GetAllReksadanaTypeOutputSchema{
			Id:            "-1",
			NamaReksadana: fmt.Sprint(err),
		}
		listType = append(listType, single)
		return listType
	}
	defer rows.Close()

	for rows.Next() {

		var id string
		var namaReksadana string
		rows.Scan(&id, &namaReksadana)
		single := response.GetAllReksadanaTypeOutputSchema{
			Id:            id,
			NamaReksadana: namaReksadana,
		}

		listType = append(listType, single)
	}

	return listType
}
