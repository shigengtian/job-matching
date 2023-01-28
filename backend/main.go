package main

import (
	"go.uber.org/zap"
	"job-maching/core"
	"job-maching/global"
	"job-maching/initialize"
)

func main() {
	global.GB_VP = core.Viper()
	initialize.OtherInit()
	global.GB_LOG = core.Zap()
	zap.ReplaceGlobals(global.GB_LOG)
	//global.GB_DB = initialize.Gorm()
	//initialize.Timer()
	//initialize.DBList()
	//if global.GB_DB != nil {
	//	initialize.RegisterTables(global.GVA_DB)
	//	// close db before program close
	//	db, _ := global.GVA_DB.DB()
	//	defer db.Close()
	//}
	core.RunWindowsServer()
}
