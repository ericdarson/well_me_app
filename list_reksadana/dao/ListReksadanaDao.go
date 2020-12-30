package dao

import (
	"fmt"
	"io/ioutil"
	"list_reksadana_api/dbconnection"
	"list_reksadana_api/entity/response"
	"os"

	_ "github.com/godror/godror"
)

type ListReksadanaDao interface {
	GetListProdukReksadana(string) []response.ListProdukReksadanaOutputSchema
}

type listReksadanaDao struct {
}

func NewListReksadana() ListReksadanaDao {
	return &listReksadanaDao{}
}
func returnErrSimluation() []response.ListProdukReksadanaOutputSchema {
	return nil
}

func (dao *listReksadanaDao) GetListProdukReksadana(idjenis string) []response.ListProdukReksadanaOutputSchema {
	var outputSchema []response.ListProdukReksadanaOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	dat, err := ioutil.ReadFile(dir + "/query/ListReksadana.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idjenis)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	defer rows.Close()

	for rows.Next() {
		var single response.ListProdukReksadanaOutputSchema
		rows.Scan(&single.ID, &single.Nama, &single.Nab, &single.Kinerja, &single.DateNab)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}
