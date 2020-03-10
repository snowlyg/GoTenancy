package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"go-tenancy/libs"
)

func NewLogFile() *os.File {
	var err error
	var f *os.File

	path := "./logs/"
	err = libs.CreateFile(path)

	filename := path + time.Now().Format("2006-01-02") + ".log"

	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		color.Red(fmt.Sprintf("日志记录出错: %v", err))
	}

	return f
}
