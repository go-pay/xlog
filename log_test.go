package xlog

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	// Cyan, "青色 debug", Reset, CyanBright, "高亮 debug", Reset, "恢复默认颜色", CyanDelLine, "删除线", Reset, CyanUnderLine, "下划线", Reset, CyanBevel, "斜体 debug", Reset, CyanBg, "背景", Reset
	// default log
	SetLevel(DebugLevel)
	SetColorful(false)

	Debug("Debug 日志")
	Info("Info 日志")
	Warn("Warn 日志")
	Error("Error 日志")

	Debugf("Debugf %s", "日志")
	Infof("Infof %s", "日志")
	Warnf("Warnf %s", "日志")
	Errorf("Errorf %s", "日志")

	log.Println("================")

	xLogger := NewLogger()
	//xLogger.Color(Red)
	xLogger.SetLevel(WarnLevel)

	xLogger.Debug("new logger Debug 日志")
	xLogger.Info("new logger Info 日志")
	xLogger.Warn("new logger Warn 日志")
	xLogger.Error("new logger Error 日志")

	xLogger.Debugf("new logger Debugf %s", "日志")
	xLogger.Infof("new logger Infof %s", "日志")
	xLogger.Warnf("new logger Warnf %s", "日志")
	xLogger.Errorf("new logger Errorf %s", "日志")

	// color log
	//Color(White).Info("color log info")
	//Color(Cyan).Debug("color log debug")
	//Color(Yellow).Warn("color log warn")
	//Color(Red).Error("color log error")
}
