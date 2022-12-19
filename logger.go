package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	TRACE_LVL = iota
	DEBUG_LVL
	INFO_LVL
	WARN_LVL
	ERROR_LVL
	FATAL_LVL
)

var Prefix string = ""
var DefaultLogger Logger

func NewLogger() Logger {
	return Logger{}
}

func ReplaceDefaultLogger(newLogger Logger) {
	DefaultLogger = newLogger
}

func TRACE(formatString string, values ...any) {
	DefaultLogger.trace(formatString, 3, values...)
}

func DEBUG(formatString string, values ...any) {
	DefaultLogger.debug(formatString, 3, values...)
}

func INFO(formatString string, values ...any) {
	DefaultLogger.info(formatString, 3, values...)
}

func WARN(formatString string, values ...any) {
	DefaultLogger.warn(formatString, 3, values...)
}

func ERROR(formatString string, values ...any) {
	DefaultLogger.error(formatString, 3, values...)
}

func FATAL(formatString string, values ...any) {
	DefaultLogger.fatal(formatString, 3, values...)
}

type Logger struct {
	traceLogger []*log.Logger
	debugLogger []*log.Logger
	warnLogger  []*log.Logger
	infoLogger  []*log.Logger
	errorLogger []*log.Logger
	fatalLogger []*log.Logger
}

func init() {
	DefaultLogger = Logger{}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	DefaultLogger.traceLogger = append(DefaultLogger.traceLogger, log.New(file, "TRACE: ", log.Ldate|log.Ltime))
	DefaultLogger.debugLogger = append(DefaultLogger.debugLogger, log.New(file, "DEBUG: ", log.Ldate|log.Ltime))
	DefaultLogger.warnLogger = append(DefaultLogger.warnLogger, log.New(file, "WARN: ", log.Ldate|log.Ltime))
	DefaultLogger.infoLogger = append(DefaultLogger.infoLogger, log.New(file, "INFO: ", log.Ldate|log.Ltime))
	DefaultLogger.errorLogger = append(DefaultLogger.errorLogger, log.New(file, "ERROR: ", log.Ldate|log.Ltime))
	DefaultLogger.errorLogger = append(DefaultLogger.fatalLogger, log.New(file, "FATAL: ", log.Ldate|log.Ltime))

	DefaultLogger.traceLogger = append(DefaultLogger.traceLogger, log.New(os.Stderr, "TRACE: ", log.Ldate|log.Ltime))
	DefaultLogger.debugLogger = append(DefaultLogger.debugLogger, log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime))
	DefaultLogger.warnLogger = append(DefaultLogger.warnLogger, log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime))
	DefaultLogger.infoLogger = append(DefaultLogger.infoLogger, log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime))
	DefaultLogger.errorLogger = append(DefaultLogger.errorLogger, log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime))
	DefaultLogger.fatalLogger = append(DefaultLogger.fatalLogger, log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime))
}

func f(lvl int) string {
	var str string = ""
	_, file, line, ok := runtime.Caller(lvl)
	//fmt.Println(runtime.CallersFrames(inp))

	/*
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		//file = short
	*/
	if ok {
		str = fmt.Sprintf("%s:%d: ", file, line)
	}
	return str
}

func (l Logger) TRACE(formatString string, values ...any) {
	l.trace(formatString, 3, values...)
}

func (l Logger) trace(formatString string, lvl int, values ...any) {

	for _, v := range l.traceLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
}

func (l Logger) DEBUG(formatString string, values ...any) {
	l.debug(formatString, 3, values...)
}

func (l Logger) debug(formatString string, lvl int, values ...any) {

	for _, v := range l.debugLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
}

//INFO is used to log to the default logger at the default level.
func (l Logger) INFO(formatString string, values ...any) {
	l.info(formatString, 3, values...)
}

func (l Logger) info(formatString string, lvl int, values ...any) {
	for _, v := range l.infoLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
}

// WARN is used to log to the default logger at the default level.
func (l Logger) WARN(formatString string, values ...any) {
	l.warn(formatString, 3, values...)
}

func (l Logger) warn(formatString string, lvl int, values ...any) {
	for _, v := range l.warnLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
}

func (l Logger) ERROR(formatString string, values ...any) {
	l.error(formatString, 3, values...)
}

func (l Logger) error(formatString string, lvl int, values ...any) {
	for _, v := range l.errorLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
}

func (l Logger) FATAL(formatString string, values ...any) {
	l.fatal(formatString, 3, values...)
}

func (l Logger) fatal(formatString string, lvl int, values ...any) {
	for _, v := range l.fatalLogger {
		v.Printf(Prefix+" "+f(lvl)+formatString+"\n", values...)
	}
	os.Exit(100)
}
