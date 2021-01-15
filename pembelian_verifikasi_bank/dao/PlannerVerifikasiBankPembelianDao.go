package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pembelian_verifikasi_bank_api/dbconnection"
	"time"

	_ "github.com/godror/godror"
)

type PlannerVerifikasiBankPembelianDao interface {
	DoVerifikasi(string, string) string
}

type plannerVerifikasiBankPembelianDao struct {
}

func NewPlannerVerifikasiBankPembelian() PlannerVerifikasiBankPembelianDao {
	return &plannerVerifikasiBankPembelianDao{}
}

func (dao *plannerVerifikasiBankPembelianDao) DoVerifikasi(vendorPw string, idTrans string) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_VERIFIKASI_BANK_BELI('%s','%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, idTrans, vendorPw)
	fmt.Println(query)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "TIME OUT"
	}

	return result
}
