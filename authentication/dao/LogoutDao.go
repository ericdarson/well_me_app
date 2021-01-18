package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type LogoutDao interface {
	Logout(token string) response.DetailLogout
}

type logoutDao struct {
}

func NewLogoutDao() LogoutDao {
	return &logoutDao{}
}
func (dao *logoutDao) Logout(token string) response.DetailLogout {
	var entityResponse response.DetailLogout
	conn := dbconnection.New()
	db := conn.GetConnection()

	var message string

	query := `BEGIN SP_LOGOUT('` + token + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &message}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return entityResponse
	} else {
		entityResponse.Message = message
	}
	return entityResponse

}
