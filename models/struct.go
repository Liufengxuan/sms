package models

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

//
type User struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Pwd  string  `json:"pwd"`
	Dept string  `json:"dept"`
	Duty string  `json:"duty"`
	Db   *sql.DB `json:"-"`
}

type Station struct {
	DbName      string `json:"dbname"`
	StationName string `json:"name"`
}
