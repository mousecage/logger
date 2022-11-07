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
)

var Logger logger

func TRACE(formatString string, values ...any) {
	Logger.trace(formatString, 3, values...)
}

func DEBUG(formatString string, values ...any) {
	Logger.DEBUG(formatString, values...)
}

func ERROR(formatString string, values ...any) {
	Logger.ERROR(formatString, values...)
}

func WARN(formatString string, values ...any) {
	Logger.WARN(formatString, values...)
}

func INFO(formatString string, values ...any) {
	Logger.INFO(formatString, values...)
}

type logger struct {
	traceLogger []*log.Logger
	debugLogger []*log.Logger
	warnLogger  []*log.Logger
	infoLogger  []*log.Logger
	errorLogger []*log.Logger
}

func init() {
	Logger = logger{}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger.traceLogger = append(Logger.traceLogger, log.New(file, "TRACE: ", log.Ldate|log.Ltime))
	Logger.debugLogger = append(Logger.debugLogger, log.New(file, "DEBUG: ", log.Ldate|log.Ltime))
	Logger.warnLogger = append(Logger.warnLogger, log.New(file, "WARN: ", log.Ldate|log.Ltime))
	Logger.infoLogger = append(Logger.infoLogger, log.New(file, "INFO: ", log.Ldate|log.Ltime))
	Logger.errorLogger = append(Logger.errorLogger, log.New(file, "ERROR: ", log.Ldate|log.Ltime))

	Logger.traceLogger = append(Logger.traceLogger, log.New(os.Stderr, "TRACE: ", log.Ldate|log.Ltime))
	Logger.debugLogger = append(Logger.debugLogger, log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime))
	Logger.warnLogger = append(Logger.warnLogger, log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime))
	Logger.infoLogger = append(Logger.infoLogger, log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime))
	Logger.errorLogger = append(Logger.errorLogger, log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime))
}

func f(lvl int) string {
	var str string = ""
	_, file, line, ok := runtime.Caller(lvl)

	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	if ok {
		str = fmt.Sprintf("%s:%d: ", file, line)
	}
	return str
}

func (l logger) TRACE(formatString string, values ...any) {
	l.trace(formatString, 3, values...)
}

func (l logger) trace(formatString string, lvl int, values ...any) {

	for _, v := range l.traceLogger {
		v.Printf(f(lvl)+formatString+"\n", values...)
	}
}

func (l logger) DEBUG(formatString string, values ...any) {
	l.debug(formatString, 3, values...)
}

func (l logger) debug(formatString string, lvl int, values ...any) {

	for _, v := range l.debugLogger {
		v.Printf(f(lvl)+formatString+"\n", values...)
	}
}

func (l logger) ERROR(formatString string, values ...any) {
	l.error(formatString, 3, values...)
}

func (l logger) error(formatString string, lvl int, values ...any) {
	for _, v := range l.errorLogger {
		v.Printf(f(lvl)+formatString+"\n", values...)
	}
}

// WARN is used to log to the default logger at the default level.
func (l logger) WARN(formatString string, values ...any) {
	l.warn(formatString, 3, values...)
}

func (l logger) warn(formatString string, lvl int, values ...any) {
	for _, v := range l.warnLogger {
		v.Printf(f(lvl)+formatString+"\n", values...)
	}
}

//INFO is used to log to the default logger at the default level.
func (l logger) INFO(formatString string, values ...any) {
	l.info(formatString, 3, values...)
}

func (l logger) info(formatString string, lvl int, values ...any) {
	for _, v := range l.infoLogger {
		v.Printf(f(lvl)+formatString+"\n", values...)
	}
}
