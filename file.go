package logger

import (
	"os"
)

type FileLog struct {
	FilePath  string
	SaveLevel Level
}

func (fileLog *FileLog) DoWrite(buf []byte) (n int, err error) {
	fd, err := os.OpenFile(fileLog.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	return fd.Write(buf)
}
