package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type GetListProdukReksadanaDao interface {
	GetListProdukReksadana(string) []response.GetListProdukReksadanaOutputSchema
}

type getListProdukReksadanaDao struct {
}

func NewListProduk() GetListProdukReksadanaDao {
	return &getListProdukReksadanaDao{}
}

func (dao *getListProdukReksadanaDao) GetListProdukReksadana(idjenis string) []response.GetListProdukReksadanaOutputSchema {
	var listProduk []response.GetListProdukReksadanaOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.GetListProdukReksadanaOutputSchema{
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			KinerjaSatuBulan: "-1",
			Nab:              "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetListProdukReksadana.query")
	if err != nil {
		fmt.Println(err)
		single := response.GetListProdukReksadanaOutputSchema{
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			KinerjaSatuBulan: "-1",
			Nab:              "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	query := string(dat)
	query = fmt.Sprintf(query, idjenis)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.GetListProdukReksadanaOutputSchema{
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			KinerjaSatuBulan: "-1",
			Nab:              "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var nama string
		var kinerja string
		var nab string
		var date string
		rows.Scan(&id, &nama, &nab, &kinerja, &date)
		single := response.GetListProdukReksadanaOutputSchema{
			Id:               id,
			Nama:             nama,
			KinerjaSatuBulan: kinerja,
			Nab:              nab,
			MaxBackwardDate:  date,
		}
		listProduk = append(listProduk, single)
	}

	return listProduk
}
