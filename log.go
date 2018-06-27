package logger

import (
	"sync"
	"log"
	//"time"
	"errors"
)

type Level uint16

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

const (
	TypeFile         = "file"
	TypeMgo          = "mongodb"
	TypeMysql        = "mysql"
)

var Logger TTLog
var logOnce sync.Once

var MysqlLogger TTLog
var mysqlLogOnce sync.Once

var MongodbLogger TTLog
var mongodbLogOnce sync.Once

type TTLog struct {
	Logger *log.Logger
}

type Log struct {
	Log DoLog
}

type FileLogConfig interface {
	GetFilePath() interface{}
}

type MysqlLogConfig interface {
	GetMysqlDns() string
	GetMysqlUser()  string
	GetMysqlPw() string
	GetMysqlDb() string
	GetMysqlTable() string
}

type MongodbLogConfig interface {
	GetMongodbDns() string
	GetMongodbPw()  string
	GetMongodbDb() string
	GetMongodbCollection() string
}

type DoLog interface {
	DoWrite(p []byte) (n int, err error)
}

/**
 * get a log driver with driver name (mysql|file|mongodb)
 * if errs
 * return nil
 */
func GetLogger(driver string, config interface{}) (*TTLog,error)  {
	if driver == TypeFile {
		if conf,ok := config.(FileLogConfig); ok {
			return GetFileLogger(conf),nil
		}else{
			return nil,errors.New("config is not impl FileLogConfig")
		}
	}

	if driver == TypeMysql {
		if conf,ok := config.(MysqlLogConfig); ok {
			return GetMysqlLogger(conf),nil
		}else{
			return nil,errors.New("config is not impl MysqlConfig")
		}
	}

	if driver == TypeMgo {
		if conf,ok := config.(MongodbLogConfig); ok {
			return GetMongodbLogger(conf),nil
		}else{
			return nil,errors.New("config is not impl MongodbLogConfig")
		}
	}

	return nil,errors.New("can't support this driver")
}

/**
 * change file path
 * if driver is file driver
 */
func (TTLog *TTLog) SetFileLogPath (config FileLogConfig){
	getFileLogger(config)
	TTLog.Logger = log.New(getFileLogger(config), "", log.LstdFlags|log.Llongfile)
}

//rewrite log Write
func (Log *Log) Write(p []byte) (n int, err error) {
	n, err = Log.Log.DoWrite(p)
	return n, err
}

//rewrite log Println
func (TTLog *TTLog) Println(v... interface{}) {
	TTLog.Logger.Println(v)
}

//get a file logger
func GetFileLogger(config FileLogConfig) (*TTLog) {
	logOnce.Do(func() {
		Logger.Logger = log.New(getFileLogger(config), "", log.LstdFlags|log.Llongfile)
	})
	return &Logger
}

//get a mysql logger
func GetMysqlLogger(config MysqlLogConfig) (*TTLog) {
	mysqlLogOnce.Do(func() {
		MysqlLogger.Logger = log.New(getMysqlLogger(config), "", log.LstdFlags|log.Llongfile)
	})
	return &MysqlLogger
}

//get a mysql logger
func GetMongodbLogger(config MongodbLogConfig) (*TTLog) {
	mongodbLogOnce.Do(func() {
		MongodbLogger.Logger = log.New(getMongodbLogger(config), "", log.LstdFlags|log.Llongfile)
	})
	return &MongodbLogger
}

//return a logger driver typeof mysql
func getMysqlLogger(config MysqlLogConfig) *Log {
	return &Log{
		Log: &MysqlLog{
			Config:config,
			mu: new(sync.Mutex),
		},
	}
}
//return a logger driver typeof file
func getFileLogger(config FileLogConfig) *Log {
	return &Log{
		Log: &FileLog{
			FilePath: config.GetFilePath().(string)+".log",
			//mu: new(sync.Mutex),
		},
	}
}
//return a logger driver typeof mongodb
func getMongodbLogger(config MongodbLogConfig) *Log {
	return &Log{
		Log: &MongodbLog{
			Config:config,
			mu: new(sync.Mutex),
		},
	}
}
