package xlog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type XLogger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)

	Debugf(fmt string, args ...any)
	Infof(fmt string, args ...any)
	Warnf(fmt string, args ...any)
	Errorf(fmt string, args ...any)
}

func (l *Logger) init() {
	l.debugLogger = log.New(os.Stdout, "[DEBUG] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	l.infoLogger = log.New(os.Stdout, "[INFO] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	l.warnLogger = log.New(os.Stdout, "[WARN] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	l.errorLogger = log.New(os.Stdout, "[ERROR] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//switch l.level {
	//case DebugLevel:
	//	l.logger = log.New(os.Stdout, "[DEBUG] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//case InfoLevel:
	//	l.logger = log.New(os.Stdout, "[INFO] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//case WarnLevel:
	//	l.logger = log.New(os.Stdout, "[WARN] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//case ErrorLevel:
	//	l.logger = log.New(os.Stdout, "[ERROR] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//default:
	//	l.logger = log.New(os.Stdout, "[INFO] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
	//}
}

type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	once        sync.Once
	level       LogLevel // 日志等级：xlog.DebugLevel、xlog.InfoLevel、xlog.WarnLevel、xlog.ErrorLevel
	colorful    bool     // 是否彩色
	col         *ColorType
}

func NewLogger() *Logger {
	return &Logger{level: InfoLevel}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Color(col *ColorType) {
	l.col = col
}

func (l *Logger) Colorful(colorful bool) {
	l.colorful = colorful
}

func (l *Logger) logout(level LogLevel, format *string, args ...any) {
	l.once.Do(func() {
		l.init()
	})
	if l.level >= level {
		switch level {
		case DebugLevel:
			l.print(l.debugLogger, CyanBright, format, args...)
		case InfoLevel:
			l.print(l.infoLogger, GreenBright, format, args...)
		case WarnLevel:
			l.print(l.warnLogger, YellowBright, format, args...)
		case ErrorLevel:
			l.print(l.errorLogger, RedBright, format, args...)
		}
	}
}

func (l *Logger) print(logger *log.Logger, col *ColorType, format *string, args ...any) {
	if l.colorful {
		if l.col != nil {
			// 设置过自定义颜色，优先自定义颜色
			if format != nil {
				_ = logger.Output(callDepth, string(*l.col)+fmt.Sprintf(*format, args...))
				return
			}
			_ = logger.Output(callDepth, string(*l.col)+fmt.Sprint(args...))
			return
		}
		// 没有自定义颜色，采用默认色系
		if format != nil {
			_ = logger.Output(callDepth, string(*col)+fmt.Sprintf(*format, args...))
			return
		}
		_ = logger.Output(callDepth, string(*col)+fmt.Sprint(args...))
		return
	}
	// 不输出颜色
	if format != nil {
		_ = logger.Output(callDepth, fmt.Sprintf(*format, args...))
		return
	}
	_ = logger.Output(callDepth, fmt.Sprintln(args...))
}

func (l *Logger) Debug(args ...any) {
	l.logout(DebugLevel, nil, args...)
}

func (l *Logger) Info(args ...any) {
	l.logout(InfoLevel, nil, args...)
}

func (l *Logger) Warn(args ...any) {
	l.logout(WarnLevel, nil, args...)
}

func (l *Logger) Error(args ...any) {
	l.logout(ErrorLevel, nil, args...)
}

func (l *Logger) Debugf(fmt string, args ...any) {
	l.logout(DebugLevel, &fmt, args...)
}

func (l *Logger) Infof(fmt string, args ...any) {
	l.logout(InfoLevel, &fmt, args...)
}

func (l *Logger) Warnf(fmt string, args ...any) {
	l.logout(WarnLevel, &fmt, args...)
}

func (l *Logger) Errorf(fmt string, args ...any) {
	l.logout(ErrorLevel, &fmt, args...)
}
