package logger

import (
	"testing"
	"fmt"
	"github.com/lwl1989/logger"
)

type configs struct {

}
/**
	GetMongodbDns() string
	GetMongodbPw()  string
	GetMongodbDb() string
	GetMongodbCollection() string

	GetMysqlDns() string
	GetMysqlUser()  string
	GetMysqlPw() string
	GetMysqlDb() string
	GetMysqlTable() string

	GetFilePath() interface{}
 */
func (configs *configs) GetMongodbDns() string {
	return "192.168.30.100:27018"
}
func (configs *configs) GetMongodbPw() string {
	return ""
}
func (configs *configs) GetMongodbDb() string {
	return "test"
}
func (configs *configs) GetMongodbCollection() string {
	return "test_coll"
}
func (configs *configs) GetMysqlDns() string {
	return "localhost:3306"
}
func (configs *configs) GetMysqlUser() string {
	return "root"
}
func (configs *configs) GetMysqlPw() string {
	return "sa"
}
func (configs *configs) GetMysqlDb() string {
	return "test"
}
func (configs *configs) GetMysqlTable() string {
	return "test_table"
}
func (configs *configs) GetFilePath() interface{} {
	return "/tmp/test"
}

func Test_GetLogger(t *testing.T) {
	var conf = &configs{}
	var drivers = []string{"mysql","mongodb","file"}
	for _,driver := range drivers {
		llog,err := logger.GetLogger(driver, conf)
		if err != nil {
			fmt.Println("mysql err")
		}else{
			llog.Println(driver)
		}
	}

	//logMysql,err := logger.GetLogger("mysql", conf)
	//if err != nil {
	//	fmt.Println("mysql err")
	//}else{
	//	logMysql.Println("mysql")
	//}
	//
	//logMongodb,err := logger.GetLogger("mongodb", conf)
	//if err != nil {
	//	fmt.Println("mongodb err")
	//}else{
	//	logMongodb.Println("mongodb")
	//}
	//
	//logFile,err := logger.GetLogger("file", conf)
	//if err != nil {
	//	fmt.Println("file err")
	//}else{
	//	logFile.Println("file")
	//}

}
