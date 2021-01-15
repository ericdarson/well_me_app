package dao

import (
	"cms_api/dbconnection"
	"cms_api/entity/request"
	"cms_api/entity/response"
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/godror/godror"
)

type BobotResikoCMSDao interface {
	InquiryBobotResiko() []response.InquiryBobotResikoOutputSchema
	UpdateBobotResiko(request.UpdateBobotResikoRequest) string
}

type bobotResikoCMSDao struct {
}

func NewBobotResikoCMS() BobotResikoCMSDao {
	return &bobotResikoCMSDao{}
}

func (dao *bobotResikoCMSDao) InquiryBobotResiko() []response.InquiryBobotResikoOutputSchema {
	var outputSchema []response.InquiryBobotResikoOutputSchema
	tempMap := make(map[string][]response.DetailBobotResiko)
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return []response.InquiryBobotResikoOutputSchema{response.InquiryBobotResikoOutputSchema{BobotResiko: "-1"}}
	}
	dat, err := ioutil.ReadFile(dir + "/query/InquiryBobotResiko.query")
	if err != nil {
		fmt.Println(err)
		return []response.InquiryBobotResikoOutputSchema{response.InquiryBobotResikoOutputSchema{BobotResiko: "-1"}}
	}
	query := string(dat)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []response.InquiryBobotResikoOutputSchema{response.InquiryBobotResikoOutputSchema{BobotResiko: "-1"}}
	}
	defer rows.Close()
	for rows.Next() {
		var tempbobot string
		var single response.DetailBobotResiko
		rows.Scan(&tempbobot, &single.Persentase, &single.IDJenis, &single.NamaJenis)

		tempMap[tempbobot] = append(tempMap[tempbobot], single)
	}

	for key, element := range tempMap {
		single := response.InquiryBobotResikoOutputSchema{
			BobotResiko:       key,
			DetailBobotResiko: element,
		}
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}

func (dao *bobotResikoCMSDao) UpdateBobotResiko(req request.UpdateBobotResikoRequest) string {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	formatquery := `BEGIN SP_UPDATE_MASTER_RESIKO(%f,%f,%d,'SYSTEM API',:1 ); END;`
	for _, i := range req.Input {
		query := fmt.Sprintf(formatquery, i.BobotResiko, i.Persentase, i.IDJenis)
		//fmt.Println(query)
		if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "ERROR"
		}
		if result != "SUKSES" {
			return result
		}
	}

	return result
}
