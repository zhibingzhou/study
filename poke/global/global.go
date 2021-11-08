package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"poke/conf"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG conf.SystemConfig
	GVA_LOG                 *zap.Logger
)
