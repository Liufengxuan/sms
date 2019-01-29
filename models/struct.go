package models

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type UserSession struct {
	IsLogin bool    `json:"isLogin"`
	User    *User   `json:"user"`
	Db      *sql.DB `json:"-"`
}

//
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"-"`
	Dept string `json:"dept"`
	Duty string `json:"duty"`
}

type Station struct {
	DbName      string `json:"dbname"`
	StationName string `json:"name"`
}
