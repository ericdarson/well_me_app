package dao

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type ProfileDao interface {
	GetById(id string) response.DetailProfile
}

type profiletDao struct {
	temp string
}

func NewProfileDao() ProfileDao {
	return &nabWidgetDao{}
}

func (dao *nabWidgetDao) GetById(id string) response.DetailProfile {
	entityResponse := response.DetailProfile{}
	conn := dbconnection.New()
	db := conn.GetConnection()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return entityResponse
	}
	dat, err := ioutil.ReadFile(dir + "\\query\\getProfile.query")
	if err != nil {
		fmt.Println(err)
		return entityResponse
	}
	query := string(dat)
	query = fmt.Sprintf(query, id)

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return entityResponse
	}
	defer rows.Close()

	for rows.Next() {

		rows.Scan(&entityResponse.BcaId, &entityResponse.Nama, &entityResponse.ProfileResiko, &entityResponse.Email, &entityResponse.SID, &entityResponse.NoRekBCA)
		//BCA_ID,NAMA,BOBOT_RESIKO,EMAIL,SID,NO_REKENING
		break

	}
	return entityResponse
}
