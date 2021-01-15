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

type ListReksadanaWidgetDao interface {
	GetReksadanaList(string, string, string) []response.ListReksadanaWidgetOutputSchema
}

type listReksadanaWidgetDao struct {
	temp string
}

func NewListReksadana() ListReksadanaWidgetDao {
	return &listReksadanaWidgetDao{}
}

func (dao *listReksadanaWidgetDao) GetReksadanaList(nama string, listKategori string, sortBy string) []response.ListReksadanaWidgetOutputSchema {
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
	dat, err := ioutil.ReadFile(dir + "/query/getListReksadana.query")
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
			filterBy = "t.nama_produk"
			orderBy = "asc"
		} else if tempSort == 2 {
			filterBy = "t.nama_produk"
			orderBy = "desc"
		} else if tempSort == 3 {
			filterBy = "round((t.nab - d.nab_daily) / d.nab_daily * 100,2)"
			orderBy = "asc"
		} else if tempSort == 4 {
			filterBy = "round((t.nab - d.nab_daily) / d.nab_daily * 100,2)"
			orderBy = "desc"
		} else if tempSort == 5 {
			filterBy = "t.nab"
			orderBy = "asc"
		} else if tempSort == 6 {
			filterBy = "t.nab"
			orderBy = "desc"
		}
	}

	query := string(dat)
	query = fmt.Sprintf(query, nama, listKategori, filterBy, orderBy)
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
		var nab string
		var kinerja string

		rows.Scan(&id, &nama, &nab, &kinerja)

		single := response.ListReksadanaWidgetOutputSchema{
			Id:   id,
			Nama: nama,
			Nab:  nab,
		}
		listProduk = append(listProduk, single)
	}

	return listProduk
}
