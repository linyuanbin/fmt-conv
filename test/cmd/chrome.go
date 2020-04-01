package main

import (
	"github.com/linyuanbin/fmt-conv/conf"
	"github.com/linyuanbin/fmt-conv/xlog"
	"github.com/linyuanbin/fmt-conv/chrome"
)

func main() {
	const op string = "main"
	config, err := conf.FromEnv()
	systemLogger := xlog.New(config.LogLevel(), "system")
	if err != nil {
		systemLogger.FatalOp(op, err)
	}
	// start Google Chrome headless.
	if err := chrome.Start(systemLogger); err != nil {
		systemLogger.FatalOp(op, err)
	}
}
