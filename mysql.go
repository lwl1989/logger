package logger

import (
	"database/sql"
	"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)


type MysqlLog struct {
	Config MysqlLogConfig
	Connection *sql.DB
	mu *sync.Mutex
}

func (mysqlLog *MysqlLog) DoWrite(buf []byte) (n int, err error) {
	mysqlLog.mu.Lock()
	defer mysqlLog.mu.Unlock()

	if mysqlLog.Connection == nil {

		err := mysqlLog.Connection.Ping()

		if err != nil {
			mysqlLog.Connection.Close()
			uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysqlLog.Config.GetMysqlUser(),
				mysqlLog.Config.GetMysqlPw(), mysqlLog.Config.GetMysqlDns(), mysqlLog.Config.GetMysqlDb())

			//必須引入匿名包mysql
			db, err := sql.Open("mysql", uri)
			mysqlLog.Connection = db

			if err != nil {
				panic("can't connection to mysql")
			}
		}

	}

	content := string(buf[:])

	if mysqlLog.Config.GetMysqlTable() != "" {
		mysqlLog.Connection.Exec(fmt.Sprintf("insert into `%s`.`%s` (`%s`) values('%s')", mysqlLog.Config.GetMysqlDb(), mysqlLog.Config.GetMysqlTable(), "content", content))
		return len(content),nil
	}

	return 0, errors.New("mysql table not do set")
}
