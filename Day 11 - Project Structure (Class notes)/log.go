package log

import "fmt"

type Logger interface {
	Info(value ...interface{})
	Error(value ...interface{})
	Warning(value ...interface{})
}

type Log struct {
}

func GetLogger() {
	return &Log{}
}

//logrus
func (l *Log) Info(value ...interface{}) {
	//operations
	fmt.Println("<<<<<<<<<<<<< INFO <<<<<<<<<<")
}
