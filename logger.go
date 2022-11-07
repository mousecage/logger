package logger

import (
	"log"
	"os"
)

const (
	TRACE = iota
	DEBUG
	ERROR
	WARN
	INFO
)

var Logger logger

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

	Logger.traceLogger = append(Logger.traceLogger, log.New(file, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.debugLogger = append(Logger.debugLogger, log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.warnLogger = append(Logger.warnLogger, log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.infoLogger = append(Logger.infoLogger, log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.errorLogger = append(Logger.errorLogger, log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile))

	Logger.traceLogger = append(Logger.traceLogger, log.New(os.Stderr, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.debugLogger = append(Logger.debugLogger, log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.warnLogger = append(Logger.warnLogger, log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.infoLogger = append(Logger.infoLogger, log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))
	Logger.errorLogger = append(Logger.errorLogger, log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile))
}

func (l logger) TRACE(formatString string, values ...any) {
	for _, v := range l.traceLogger {
		v.Printf(formatString, values...)
	}
}

func (l logger) DEBUG(formatString string, values ...any) {
	for _, v := range l.debugLogger {
		v.Printf(formatString, values...)
	}
}

func (l logger) ERROR(formatString string, values ...any) {
	for _, v := range l.errorLogger {
		v.Printf(formatString, values...)
	}
}

func (l logger) WARN(formatString string, values ...any) {
	for _, v := range l.warnLogger {
		v.Printf(formatString, values...)
	}
}

func (l logger) INFO(formatString string, values ...any) {
	for _, v := range l.infoLogger {
		v.Printf(formatString, values...)
	}
}
