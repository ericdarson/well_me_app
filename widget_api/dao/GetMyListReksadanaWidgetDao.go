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

type MyListReksadanaWidgetDao interface {
	GetMyReksadanaList(string, string, string, string) []response.ListReksadanaWidgetOutputSchema
}

type myListReksadanaWidgetDao struct {
	temp string
}

func NewMyListReksadana() MyListReksadanaWidgetDao {
	return &myListReksadanaWidgetDao{}
}

func (dao *myListReksadanaWidgetDao) GetMyReksadanaList(bcaid string, nama string, listKategori string, sortBy string) []response.ListReksadanaWidgetOutputSchema {
	var listProduk []response.ListReksadanaWidgetOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.ListReksadanaWidgetOutputSchema{
			Id: "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	dat, err := ioutil.ReadFile(dir + "/query/getMyReksadanaList.query")
	if err != nil {
		fmt.Println(err)
		single := response.ListReksadanaWidgetOutputSchema{
			Id: "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
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
			filterBy = "p.nama_produk"
			orderBy = "asc"
		} else if tempSort == 2 {
			filterBy = "p.nama_produk"
			orderBy = "desc"
		} else if tempSort == 7 {
			filterBy = "(sum(r.NAB_sekarang * r.jumlah_unit) - sum(r.NAB_Rerata * r.jumlah_unit))"
			orderBy = "asc"
		} else if tempSort == 8 {
			filterBy = "(sum(r.NAB_sekarang * r.jumlah_unit) - sum(r.NAB_Rerata * r.jumlah_unit))"
			orderBy = "desc"
		}
	}

	query := string(dat)
	query = fmt.Sprintf(query, bcaid, nama, listKategori, filterBy, orderBy)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.ListReksadanaWidgetOutputSchema{
			Id: "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var nama string
		var profit string

		rows.Scan(&id, &nama, &profit)

		single := response.ListReksadanaWidgetOutputSchema{
			Id:     id,
			Nama:   nama,
			Profit: profit,
		}
		listProduk = append(listProduk, single)
	}

	return listProduk
}
