package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"job-maching/config"
	"sync"
)

var (
	GB_DB     *gorm.DB
	GB_DBList map[string]*gorm.DB
	GB_REDIS  *redis.Client
	GB_CONFIG config.Server
	GB_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GB_LOG *zap.Logger
	//GB_Timer               timer.Timer = timer.NewTimerTask()
	//GB_Concurrency_Control = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)

func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GB_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GB_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
