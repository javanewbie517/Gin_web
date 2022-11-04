package global

import (
	"bubble/server/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	lock       sync.RWMutex
)
