package clog

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/clog/logger"
	"log"
	"path"
	"runtime"
	"strconv"
)

type Clog struct{}

var (
	ClogItem *Clog
)

var ClogGet = func() *Clog {
	if ClogItem == nil {
		ClogItem = &Clog{}
		return ClogItem
	} else {
		return ClogItem
	}
}

func Println(str interface{}) {
	logger.Reset()
	logger.SetTimeFormat("2006-01-02 15.04.05")
	logger.SetLevel(logger.LevelDebug)
	logger.SetColorMod(true) // 开启颜色打印
	//logger.SetLocation(1) // 打印调用方法的位置
	logger.SetLocation(2) // 打印调用文件的位置
	sprintf := fmt.Sprintf("--> ss%v", str)
	bg := logger.Green(sprintf)
	logger.PInfo(bg)
}

func PrintPg(str interface{}) {
	logger.Reset()
	logger.SetTimeFormat("2006-01-02 15.04.05")
	logger.SetLevel(logger.LevelDebug)
	logger.SetColorMod(true) // 开启颜色打印
	//logger.SetLocation(1) // 打印调用方法的位置
	logger.SetLocation(2) // 打印调用文件的位置
	sprintf := fmt.Sprintf("--> ss%v", str)
	bg := logger.Red(sprintf)
	logger.PDebug(bg)
}

func PrintWa(str interface{}) {
	logger.Reset()
	logger.SetTimeFormat("2006-01-02 15.04.05")
	logger.SetLevel(logger.LevelDebug)
	logger.SetColorMod(true) // 开启颜色打印
	//logger.SetLocation(1) // 打印调用方法的位置
	logger.SetLocation(2) // 打印调用文件的位置
	sprintf := fmt.Sprintf("--> ss%v", str)
	bg := logger.Red(sprintf)
	logger.PWarn(bg)
}


func Sprint(str string) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "[" + filename + ":" + strconv.Itoa(line) + "] " + str
	return msg
}

func (c *Clog) logFile(str string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "[" + filename + ":" + strconv.Itoa(line) + "] " + str

	log.Println(msg)
}
