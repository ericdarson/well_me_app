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

type DailyNabCMSDao interface {
	InquiryDailyNab(string) []response.DailyNabOutputSchema
	InsertDailyNab(request.InsertDailyProductsRequest) response.InsertDailyProductsOutputSchema
}

type dailyNabCMSDao struct {
}

func NewDailyNabCMS() DailyNabCMSDao {
	return &dailyNabCMSDao{}
}

func returnErrDailyNab() []response.DailyNabOutputSchema {
	return []response.DailyNabOutputSchema{response.DailyNabOutputSchema{IDProduk: -1}}
}

func (dao *dailyNabCMSDao) InquiryDailyNab(idJenis string) []response.DailyNabOutputSchema {
	var outputSchema []response.DailyNabOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrDailyNab()
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetDailyProducts.query")
	if err != nil {
		fmt.Println(err)
		return returnErrDailyNab()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idJenis)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrDailyNab()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.DailyNabOutputSchema{}
		rows.Scan(&single.IDProduk, &single.NamaProduk, &single.Nab, &single.IsUpdatedToday)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}

func (dao *dailyNabCMSDao) InsertDailyNab(req request.InsertDailyProductsRequest) response.InsertDailyProductsOutputSchema {
	var outputSchema response.InsertDailyProductsOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	for _, iter := range req.Input {
		var result string
		query := `BEGIN SP_UPDATE_PRODUK_DAILY(%d,%f,'SYSTEM API',:1 ); END;`
		query = fmt.Sprintf(query, iter.IDProduk, iter.Nab)
		if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return response.InsertDailyProductsOutputSchema{ListGagal: []request.InsertDailyProduct{request.InsertDailyProduct{IDProduk: -1}}}
		}
		if result == "GAGAL" {
			outputSchema.ListGagal = append(outputSchema.ListGagal, iter)
		}
	}
	return outputSchema
}
