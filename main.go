package main

import (
	"github.com/sh-tatsuno/shabm/config"
	"github.com/sh-tatsuno/shabm/controller"
	"github.com/sh-tatsuno/shabm/util"
)

func main() {
	util.LoggingSettings(config.Config.LogFile)
	controller.WebServer()
}
