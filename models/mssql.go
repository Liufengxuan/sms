package models

import (
	"bytes"
	"database/sql"

	"github.com/astaxie/beego/config"
	_ "github.com/denisenkom/go-mssqldb"
)

//	db, err := sql.Open("mssql", "server=ALPHA\\SQLEXPRESS;port=1433;database=CPMS;user id=sa;password=8139;encrypt=disable ")

//GetDB 通过给定数据库名字获取数据库对象
func GetDB(dbName string) (*sql.DB, error) {
	ini, err := config.NewConfig("ini", "conf/db.conf")
	//test//	ini, err := config.NewConfig("ini", "D:/WorkSpace/MyGo/src/sms/conf/db.conf")
	if err != nil {
		return nil, err
	}

	dbServer := ini.String("sqlserver::server")
	dbPort := ini.String("sqlserver::port")
	dbUserId := ini.String("sqlserver::userid")
	dbPassword := ini.String("sqlserver::password")
	dbEncrypt := ini.String("sqlserver::encrypt")

	var buf bytes.Buffer
	if dbServer != "" {
		buf.WriteString("server=")
		buf.WriteString(dbServer + ";")
	}
	if dbPort != "" {
		buf.WriteString("port=")
		buf.WriteString(dbPort + ";")
	}
	if dbUserId != "" {
		buf.WriteString("user id=")
		buf.WriteString(dbUserId + ";")
	}
	if dbName != "" {
		buf.WriteString("database=")
		buf.WriteString(dbName + ";")
	}

	if dbPassword != "" {
		buf.WriteString("password=")
		buf.WriteString(dbPassword + ";")
	}
	if dbEncrypt != "" {
		buf.WriteString("encrypt=")
		buf.WriteString(dbEncrypt)
	}
	db, err := sql.Open("mssql", buf.String())
	if err != nil {
		return nil, err
	} else {
		err = db.Ping()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
