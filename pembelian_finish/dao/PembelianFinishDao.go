package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pembelian_finish_api/dbconnection"
	"time"

	_ "github.com/godror/godror"
)

type PembelianFinishDao interface {
	DoneBeli(string, string) string
}

type pembelianFinishDao struct {
}

func NewPembelianFinish() PembelianFinishDao {
	return &pembelianFinishDao{}
}

func (dao *pembelianFinishDao) DoneBeli(vendorPw string, idTrans string) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_VERIFIKASI_PEMBELIAN('%s','%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, idTrans, vendorPw)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "TIME OUT"
	}

	return result
}
