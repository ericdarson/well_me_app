package dao

import (
	"database/sql"
	"log"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type LogoutAdminDao interface {
	LogoutAdmin(token string) response.DetailLogout
}

type logoutAdminDao struct {
}

func NewLogoutAdminDao() LogoutAdminDao {
	return &logoutAdminDao{}
}
func (dao *logoutAdminDao) LogoutAdmin(token string) response.DetailLogout {
	var entityResponse response.DetailLogout
	conn := dbconnection.New()
	db := conn.GetConnection()

	var message string

	query := `BEGIN SP_LOGOUT_ADMIN('` + token + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &message}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return entityResponse
	} else {
		entityResponse.Message = message
	}
	return entityResponse

}
