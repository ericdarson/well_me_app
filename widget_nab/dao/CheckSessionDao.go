package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type CheckSessionDao interface {
	CheckSession(bcaId string, token string) response.DetailCheckSession
}

type checkSessionDao struct {
}

func NewCheckSessionDao() CheckSessionDao {
	return &checkSessionDao{}
}

func (dao *checkSessionDao) CheckSession(bcaId string, token string) response.DetailCheckSession {
	conn := dbconnection.New()
	db := conn.GetConnection()
	entityResponse := response.DetailCheckSession{}
	query := `BEGIN SP_Update_Session('` + bcaId + `','` + token + `',:1,:2); END;`

	if _, err := db.Exec(query, sql.Out{Dest: &entityResponse.Message}, sql.Out{Dest: &entityResponse.Token}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return entityResponse
	} else {

		return entityResponse
	}
}
