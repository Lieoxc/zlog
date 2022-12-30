package main

import (
	"fmt"

	"github.com/Lieoxc/zlog"
	"go.uber.org/zap"
)

func main() {
	lg := zlog.NewLogger(zlog.SetAppName("test_app"), zlog.SetDevelopment(true),
		zlog.SetLevel(zap.DebugLevel))

	lg.Debug(fmt.Sprint("debug log ", 1), zap.Int("line", 47))
	lg.Info(fmt.Sprint("Info log ", 2), zap.Any("level", "1231231231"))
	lg.Warn(fmt.Sprint("warn log ", 3), zap.String("level", `{"a":"4","b":"5"}`))
	lg.Error(fmt.Sprint("err log ", 4), zap.String("level", `{"a":"7","b":"8"}`))

}
