package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type GetJenisReksanaDanaDao interface {
	GetJenisReksadana() []response.GetJenisReksadanaOutputSchema
}

type getJenisReksanaDanaDao struct {
}

func New() GetJenisReksanaDanaDao {
	return &getJenisReksanaDanaDao{}
}

func (dao *getJenisReksanaDanaDao) GetJenisReksadana() []response.GetJenisReksadanaOutputSchema {
	var listReksadana []response.GetJenisReksadanaOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.GetJenisReksadanaOutputSchema{
			Id:        "-1",
			Reksadana: fmt.Sprint(err),
		}
		listReksadana = append(listReksadana, single)
		return listReksadana
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetJenisReksadana.query")
	if err != nil {
		fmt.Println(err)
		single := response.GetJenisReksadanaOutputSchema{
			Id:        "-1",
			Reksadana: fmt.Sprint(err),
		}
		listReksadana = append(listReksadana, single)
		return listReksadana
	}
	query := string(dat)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.GetJenisReksadanaOutputSchema{
			Id:        "-1",
			Reksadana: fmt.Sprint(err),
		}
		listReksadana = append(listReksadana, single)
		return listReksadana
	}
	defer rows.Close()

	for rows.Next() {

		var reksadana string
		var id string
		rows.Scan(&id, &reksadana)
		single := response.GetJenisReksadanaOutputSchema{
			Id:        id,
			Reksadana: reksadana,
		}
		listReksadana = append(listReksadana, single)
	}

	return listReksadana
}
