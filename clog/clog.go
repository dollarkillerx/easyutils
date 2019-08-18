package clog

import (
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

func LogFunc(str string) {
	ClogGet().logFunc(str)
}

func LogFile(str string) {
	ClogGet().logFile(str)
}

func (c *Clog) logFunc(str string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "[" + filename + ":" + strconv.Itoa(line) + "] " + str

	log.Print(msg)
}

func (c *Clog) logFile(str string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "[" + filename + ":" + strconv.Itoa(line) + "] " + str

	log.Print(msg)
}
