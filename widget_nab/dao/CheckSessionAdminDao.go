package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type CheckSessionAdminDao interface {
	CheckSessionAdmin(username string, token string) response.DetailCheckAdminSession
}

type checkSessionAdminDao struct {
}

func NewCheckSessionAdminDao() CheckSessionAdminDao {
	return &checkSessionAdminDao{}
}

func (dao *checkSessionAdminDao) CheckSessionAdmin(username string, token string) response.DetailCheckAdminSession {
	conn := dbconnection.New()
	db := conn.GetConnection()
	entityResponse := response.DetailCheckAdminSession{}
	query := `BEGIN SP_UPDATE_SESSION_ADMIN('` + username + `','` + token + `',:1,:2); END;`

	if _, err := db.Exec(query, sql.Out{Dest: &entityResponse.Message}, sql.Out{Dest: &entityResponse.Token}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		entityResponse.Username = username
		return entityResponse
	} else {
		entityResponse.Username = username
		return entityResponse
	}
}
