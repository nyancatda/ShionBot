package utils

import (
	"fmt"
	"os"
	"time"
)

//写入Log进文件
func LogWrite(LogText string) {
	_, err := os.Stat("./log")
	if err != nil {
		os.MkdirAll("./log", 0777)
	}
	now := time.Now()
	logFileName := "log/" + now.Format("2006-01-02") + ".log"
	f, err := os.OpenFile(logFileName, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write([]byte(LogText + "\r\n"))
	if err != nil {
		fmt.Println(err)
	}
}
