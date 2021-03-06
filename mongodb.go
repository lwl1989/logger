package logger

import (
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongodbLog struct {
	Config MongodbLogConfig
	Connection *mgo.Session
//	mu *sync.Mutex
}

type mongodbLogContent struct {
	Content string `json:"content"`
	CreateTime bson.MongoTimestamp `json:"create_time"`
} 

func (mongodbLog *MongodbLog) DoWrite(buf []byte) (n int, err error) {
	//mongodbLog.mu.Lock()
	//defer mongodbLog.mu.Unlock()

	if mongodbLog.Connection == nil {
		mongodbLog.DoConnection()
	}else{
		err = mongodbLog.Connection.Ping()
		if err != nil {
			mongodbLog.Connection.Close()
			mongodbLog.DoConnection()
		}
	}

	t := bson.MongoTimestamp(time.Now().UnixNano())
	content := &mongodbLogContent{
		Content:string(buf[:]),
		CreateTime:t,
	}

	err = mongodbLog.Connection.DB(mongodbLog.Config.GetMongodbDb()).
		C(mongodbLog.Config.GetMongodbCollection()).Insert(content)

	if err != nil {
		return 0, errors.New("mongodb write error with content" + content.Content)
	}

	return len(content.Content), nil
}

func (mongodbLog *MongodbLog) DoConnection() {
	var err error
	mongodbLog.Connection, err = mgo.Dial(mongodbLog.Config.GetMongodbDns())
	if err != nil {
		panic("mongodb can't connection")
	}
}

