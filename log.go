package xlog

const (
	ErrorLevel LogLevel = iota + 1
	WarnLevel
	InfoLevel
	DebugLevel
)

const (
	callDepth = 4
)

type LogLevel int

var (
	logger = &Logger{level: InfoLevel}
)

func SetLevel(level LogLevel) {
	logger.level = level
}

func SetColorful(colorful bool) {
	logger.colorful = colorful
}

// 青色:Cyan
func Debug(args ...any) {
	logger.logout(DebugLevel, nil, args...)
}

// 青色:Cyan
func Debugf(format string, args ...any) {
	logger.logout(DebugLevel, &format, args...)
}

// 白色:White
func Info(args ...any) {
	logger.logout(InfoLevel, nil, args...)
}

// 白色:White
func Infof(format string, args ...any) {
	logger.logout(InfoLevel, &format, args...)
}

// 黄色:Yellow
func Warn(args ...any) {
	logger.logout(WarnLevel, nil, args...)
}

// 黄色:Yellow
func Warnf(format string, args ...any) {
	logger.logout(WarnLevel, &format, args...)
}

// 红色:Red
func Error(args ...any) {
	logger.logout(ErrorLevel, nil, args...)
}

// 红色:Red
func Errorf(format string, args ...any) {
	logger.logout(ErrorLevel, &format, args...)
}
