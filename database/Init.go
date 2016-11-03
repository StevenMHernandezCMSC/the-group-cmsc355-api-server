package database

import (
	"os"
	"fmt"
	"github.com/coopernurse/gorp"
	"database/sql"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	if os.Getenv("DBUSER") == "" {
		b, _ := ioutil.ReadFile("./config/database.yml")

		obj := DBConfig{}
		yaml.Unmarshal(b, &obj)

		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", obj.User, obj.Password, obj.Host, obj.Port, obj.DBname)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBTABLE"))
}
