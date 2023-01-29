package main

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"job-maching/core"
	"job-maching/global"
	"job-maching/initialize"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)

	dsn := "host=localhost user=shi password=s#nd0#B9@iW5 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error %s", err.Error())
	} else {
		global.DB = db
	}

	if global.DB != nil {
		initialize.RegisterTables(global.DB)
		// close db before program close
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
