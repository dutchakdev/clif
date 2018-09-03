package db

import (
	"database/sql"
	"github.com/dutchakdev/clif/helpers"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/user"
)

func DbConnection() (*sql.DB, error) {
	u, _ := user.Current()
	var path = u.HomeDir + "/.clif/";

	os.MkdirAll(path, os.ModePerm)
	return sql.Open("sqlite3", path + "clif.db")
}

func PrepareDatabase()  {
	database, err := DbConnection()
	helpers.CheckErr(err)
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY, name TEXT)")
	helpers.CheckErr(err)
	statement.Exec()
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, name TEXT, command TEXT, category INTEGER)")
	helpers.CheckErr(err)
	statement.Exec()
}