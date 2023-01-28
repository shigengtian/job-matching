package core

import (
	"fmt"
	"go.uber.org/zap"
	"job-maching/global"
	"job-maching/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
	//	initialize.Redis()
	//}

	//if global.GB_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GB_CONFIG.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GB_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	Buy me coffe
`, address)
	global.GB_LOG.Error(s.ListenAndServe().Error())
}
