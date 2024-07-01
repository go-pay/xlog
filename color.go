package xlog

type ColorType string

var (
	Reset = &reset
	// 标准
	White   = &white
	Red     = &red
	Green   = &green
	Yellow  = &yellow
	Blue    = &blue
	Magenta = &magenta
	Cyan    = &cyan
	// 高亮
	WhiteBright   = &whiteBright
	RedBright     = &redBright
	GreenBright   = &greenBright
	YellowBright  = &yellowBright
	BlueBright    = &blueBright
	MagentaBright = &magentaBright
	CyanBright    = &cyanBright
	// 斜体
	WhiteBevel   = &whiteBevel
	RedBevel     = &redBevel
	GreenBevel   = &greenBevel
	YellowBevel  = &yellowBevel
	BlueBevel    = &blueBevel
	MagentaBevel = &magentaBevel
	CyanBevel    = &cyanBevel
	// 下划线
	WhiteUnderLine   = &whiteUnderLine
	RedUnderLine     = &redUnderLine
	GreenUnderLine   = &greenUnderLine
	YellowUnderLine  = &yellowUnderLine
	BlueUnderLine    = &blueUnderLine
	MagentaUnderLine = &magentaUnderLine
	CyanUnderLine    = &cyanUnderLine
	// 背景色
	WhiteBg   = &whiteBg
	RedBg     = &redBg
	GreenBg   = &greenBg
	YellowBg  = &yellowBg
	BlueBg    = &blueBg
	MagentaBg = &magentaBg
	CyanBg    = &cyanBg
	// 删除线
	WhiteDelLine   = &whiteDelLine
	RedDelLine     = &redDelLine
	GreenDelLine   = &greenDelLine
	YellowDelLine  = &yellowDelLine
	BlueDelLine    = &blueDelLine
	MagentaDelLine = &magentaDelLine
	CyanDelLine    = &cyanDelLine

	// ================================================

	reset = ColorType([]byte{27, 91, 48, 109})
	// 标准
	white   = ColorType([]byte{27, 91, 51, 48, 109}) // 白色
	red     = ColorType([]byte{27, 91, 51, 49, 109}) // 红色
	green   = ColorType([]byte{27, 91, 51, 50, 109}) // 绿色
	yellow  = ColorType([]byte{27, 91, 51, 51, 109}) // 黄色
	blue    = ColorType([]byte{27, 91, 51, 52, 109}) // 蓝色
	magenta = ColorType([]byte{27, 91, 51, 53, 109}) // 紫色
	cyan    = ColorType([]byte{27, 91, 51, 54, 109}) // 青色
	// 高亮
	whiteBright   = ColorType([]byte{27, 91, 49, 59, 51, 48, 109})
	redBright     = ColorType([]byte{27, 91, 49, 59, 51, 49, 109})
	greenBright   = ColorType([]byte{27, 91, 49, 59, 51, 50, 109})
	yellowBright  = ColorType([]byte{27, 91, 49, 59, 51, 51, 109})
	blueBright    = ColorType([]byte{27, 91, 49, 59, 51, 52, 109})
	magentaBright = ColorType([]byte{27, 91, 49, 59, 51, 53, 109})
	cyanBright    = ColorType([]byte{27, 91, 49, 59, 51, 54, 109})
	// 斜体
	whiteBevel   = ColorType([]byte{27, 91, 51, 59, 51, 48, 109})
	redBevel     = ColorType([]byte{27, 91, 51, 59, 51, 49, 109})
	greenBevel   = ColorType([]byte{27, 91, 51, 59, 51, 50, 109})
	yellowBevel  = ColorType([]byte{27, 91, 51, 59, 51, 51, 109})
	blueBevel    = ColorType([]byte{27, 91, 51, 59, 51, 52, 109})
	magentaBevel = ColorType([]byte{27, 91, 51, 59, 51, 53, 109})
	cyanBevel    = ColorType([]byte{27, 91, 51, 59, 51, 54, 109})
	// 下划线
	whiteUnderLine   = ColorType([]byte{27, 91, 52, 59, 51, 48, 109})
	redUnderLine     = ColorType([]byte{27, 91, 52, 59, 51, 49, 109})
	greenUnderLine   = ColorType([]byte{27, 91, 52, 59, 51, 50, 109})
	yellowUnderLine  = ColorType([]byte{27, 91, 52, 59, 51, 51, 109})
	blueUnderLine    = ColorType([]byte{27, 91, 52, 59, 51, 52, 109})
	magentaUnderLine = ColorType([]byte{27, 91, 52, 59, 51, 53, 109})
	cyanUnderLine    = ColorType([]byte{27, 91, 52, 59, 51, 54, 109})
	// 背景色
	whiteBg   = ColorType([]byte{27, 91, 55, 59, 51, 48, 109})
	redBg     = ColorType([]byte{27, 91, 55, 59, 51, 49, 109})
	greenBg   = ColorType([]byte{27, 91, 55, 59, 51, 50, 109})
	yellowBg  = ColorType([]byte{27, 91, 55, 59, 51, 51, 109})
	blueBg    = ColorType([]byte{27, 91, 55, 59, 51, 52, 109})
	magentaBg = ColorType([]byte{27, 91, 55, 59, 51, 53, 109})
	cyanBg    = ColorType([]byte{27, 91, 55, 59, 51, 54, 109})
	// 删除线
	whiteDelLine   = ColorType([]byte{27, 91, 57, 59, 51, 48, 109})
	redDelLine     = ColorType([]byte{27, 91, 57, 59, 51, 49, 109})
	greenDelLine   = ColorType([]byte{27, 91, 57, 59, 51, 50, 109})
	yellowDelLine  = ColorType([]byte{27, 91, 57, 59, 51, 51, 109})
	blueDelLine    = ColorType([]byte{27, 91, 57, 59, 51, 52, 109})
	magentaDelLine = ColorType([]byte{27, 91, 57, 59, 51, 53, 109})
	cyanDelLine    = ColorType([]byte{27, 91, 57, 59, 51, 54, 109})
)

//var cl *ColorLogger

//type ColorLogger struct {
//	Color *ColorType
//}
//
//func Color(color *ColorType) *ColorLogger {
//	if cl == nil {
//		cl = &ColorLogger{
//			Color: color,
//		}
//		return cl
//	}
//	cl.Color = color
//	return cl
//}
//
//func (l *ColorLogger) Info(args ...any) {
//	l.i.LogOut(l.Color, nil, args...)
//}
//
//func (l *ColorLogger) Infof(format string, args ...any) {
//	l.i.LogOut(l.Color, &format, args...)
//}
//
//func (l *ColorLogger) Debug(args ...any) {
//	l.d.LogOut(l.Color, nil, args...)
//}
//
//func (l *ColorLogger) Debugf(format string, args ...any) {
//	l.d.LogOut(l.Color, &format, args...)
//}
//
//func (l *ColorLogger) Warn(args ...any) {
//	l.w.LogOut(l.Color, nil, args...)
//}
//
//func (l *ColorLogger) Warnf(format string, args ...any) {
//	l.w.LogOut(l.Color, &format, args...)
//}
//
//func (l *ColorLogger) Error(args ...any) {
//	l.e.LogOut(l.Color, nil, args...)
//}
//
//func (l *ColorLogger) Errorf(format string, args ...any) {
//	l.e.LogOut(l.Color, &format, args...)
//}
