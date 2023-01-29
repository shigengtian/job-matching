package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"job-maching/global"
	"job-maching/model/master"
	"job-maching/model/system"
	"os"
)

// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// システム
		system.SysApi{},
		//master.Authority{},
		//master.Industry{},
		master.Occupation{},
		//master.Menu{},
		master.ProgramLanguage{},
		//master.User{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
