package database

import (
	"os"
	"fmt"
	"github.com/coopernurse/gorp"
	"database/sql"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/utils"
)

func Init() gorp.DbMap {
	db, err := sql.Open("mysql", getDBConnectionString());
	if (err != nil) {
		utils.LogError(err, "sql.Open failed")
	}

	dbmap := gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap

}

func getDBConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBTABLE"))
}
