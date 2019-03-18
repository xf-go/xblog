package main

import (
	"flag"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"

	"xblog/system"
	"xblog/model"
)

func main() {
	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")

	// flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		seelog.Critical("err parsing seelog config file: ", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	if err := system.LoadConfiguration(*configFilePath); err != nil {
		seelog.Critical("err parsing config log file: ", err)
		return
	}

	db, err := model.InitDB()
	if err != nil {
		seelog.Critical("err open databases: ", err)
		return
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
}
