package log

import "fmt"

type Logger interface {
	Info(value ...interface{})
	Error(value ...interface{})
	Warning(value ...interface{})
}
type Log struct {
}

//logrus
func (l *Log) Info(value ...interface{}) {
	//operations...
	fmt.Println("<<<<<<<<<<<INFO<<<<<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<<<<<<<INFO<<<<<<<<<")
}
func (l *Log) Error(value ...interface{}) {
	fmt.Println("<<<<<<<<<<<Error<<<<<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<<<<<<<Error<<<<<<<<<")
}
func (l *Log) Warning(value ...interface{}) {
	fmt.Println("<<<<<<<<<<<Warning<<<<<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<<<<<<<Warning<<<<<<<<<")
}
func GetLogger() Logger {
	return &Log{}
}
