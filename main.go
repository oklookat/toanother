package main

import (
	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/datadir"
	"github.com/oklookat/toanother/core/logger"
	"github.com/oklookat/toanother/ui"
)

func main() {
	var err error

	// datadir.
	if err = datadir.Init(); err != nil {
		panic("[datadir] " + err.Error())
	}

	// config.
	if err = base.ConfigFile.Init(); err != nil {
		panic("[config] " + err.Error())
	}

	// logger.
	if err = logger.Init(); err != nil {
		panic("[logger] " + err.Error())
	}

	// ui.
	var ui = ui.Instance{}
	ui.Start()
}
