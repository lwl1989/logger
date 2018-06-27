package logger

import (
	"os"
	//"sync"
	"time"
	"fmt"
)

type FileLog struct {
	FilePath  string
	SaveLevel Level
	//mu *sync.Mutex
}



//func IsDir(dirname string) bool  {
//	handler, err := os.Stat(dirname)
//	if ! (err == nil || os.IsExist(err))  {
//		return false
//	}else {
//		return handler.IsDir()
//	}
//}
//
//func IsFile(filename string) bool  {
//	handler, err := os.Stat(filename);
//	if ! (err == nil || os.IsExist(err))  {
//		return false
//	}else if handler.IsDir() {
//		return false
//	}
//	return true
//}
//
//
//func IsExist(path string) (bool, error){
//	_, err := os.Stat(path)
//
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
//
//func GetFileByteSize(filename string) (bool,int64) {
//	if ! IsFile(filename) {
//		return false,0
//	}
//	handler, _ := os.Stat(filename)
//	return true, handler.Size()
//}


func (fileLog *FileLog) DoWrite(buf []byte) (n int, err error) {
	//fileLog.mu.Lock()
	//defer fileLog.mu.Unlock()


	filePath := fmt.Sprintf("%s_%s.log", fileLog.FilePath, time.Now().Format("2006-01-02"))
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return 0, err
	}
	defer fd.Close()

	return fd.Write(buf)
}
