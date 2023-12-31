package xlog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type WarnLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (w *WarnLogger) LogOut(col *ColorType, format *string, v ...any) {
	w.once.Do(func() {
		w.init()
	})
	if Level >= WarnLevel {
		if col != nil {
			if format != nil {
				_ = w.logger.Output(callDepth, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			_ = w.logger.Output(callDepth, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			_ = w.logger.Output(callDepth, fmt.Sprintf(*format, v...))
			return
		}
		_ = w.logger.Output(callDepth, fmt.Sprintln(v...))
	}
}

func (w *WarnLogger) init() {
	if Level == 0 {
		Level = WarnLevel
	}
	w.logger = log.New(os.Stdout, "[WARN] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
