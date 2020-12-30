package dbconnection

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/godror/godror"
)

type Connection interface {
	GetConnection() *sql.DB
}

var conn Connection

type connection struct {
	db *sql.DB
}

func init() {
	conn = &connection{
		db: nil,
	}
}

func New() Connection {
	return conn
}

func (conn *connection) GetConnection() *sql.DB {
	if conn.db == nil {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		dat, err := ioutil.ReadFile(dir + "/properties/connectionstring.properties")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		db, err := sql.Open("godror", string(dat))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		conn.db = db
		//defer conn.db.Close()
		return db
	}
	return conn.db
}
