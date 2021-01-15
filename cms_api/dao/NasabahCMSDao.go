package dao

import (
	"cms_api/dbconnection"
	"cms_api/entity/response"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/godror/godror"
)

type NasabahCMSDao interface {
	InquiryDataNasabah(string) response.InquiryDataNasabahOutputSchema
}

type nasabahCMSDao struct {
}

func NewNasabahCMS() NasabahCMSDao {
	return &nasabahCMSDao{}
}

func returnErrNasabah() response.InquiryDataNasabahOutputSchema {
	return response.InquiryDataNasabahOutputSchema{Nama: "-1"}
}

func (dao *nasabahCMSDao) InquiryDataNasabah(bcaid string) response.InquiryDataNasabahOutputSchema {
	var outputSchema response.InquiryDataNasabahOutputSchema

	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrNasabah()
	}
	dat, err := ioutil.ReadFile(dir + "/query/InquiryDataNasabah.query")
	if err != nil {
		fmt.Println(err)
		return returnErrNasabah()
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaid)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrNasabah()
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&outputSchema.Nama, &outputSchema.Email, &outputSchema.NoRekening, &outputSchema.SID, &outputSchema.TglJoin, &outputSchema.WrongAttempt, &outputSchema.DateLocked, &outputSchema.BobotResiko, &outputSchema.LevelResiko, &outputSchema.TingkatResiko, &outputSchema.NoHP)
	}

	return outputSchema
}
