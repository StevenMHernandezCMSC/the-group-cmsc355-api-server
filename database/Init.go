package database

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"database/sql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
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
	b, _ := ioutil.ReadFile("./app/config/database.yml")

	obj := DBConfig{}
	yaml.Unmarshal(b, &obj)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", obj.User, obj.Password, obj.Host, obj.Port, obj.DBname)
}
