package logger

import (
	"database/sql"
	"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)



type MysqlLog struct {
	DB  string
	TABLE string
	Config MysqlLogConfig
	Connection *sql.DB
}

func (mysqlLog *MysqlLog) DoWrite(buf []byte) (n int, err error) {
	if mysqlLog.Connection == nil {
		uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",mysqlLog.Config.GetMysqlUser(),
			mysqlLog.Config.GetMysqlPw(),mysqlLog.Config.GetMysqlDns(),mysqlLog.Config.GetMysqlDb())

		//必須引入匿名包mysql
		db, err := sql.Open("mysql", uri)
		if err != nil {
			panic("can't connection to mysql")
		}
		mysqlLog.Connection = db
	}

	content := string(buf[:])

	if mysqlLog.TABLE != "" {
		mysqlLog.Connection.Exec(fmt.Sprintf("insert into `%s`.`%s` (`%s`) values('%s')", mysqlLog.DB, mysqlLog.TABLE, "content", content))
		return len(content),nil
	}

	return 0, errors.New("mysql table not do set")
}
