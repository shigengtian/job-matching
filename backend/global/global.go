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
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	LOG *zap.Logger
	//GB_Timer               timer.Timer = timer.NewTimerTask()
	//GB_Concurrency_Control = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)
