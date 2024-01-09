package core

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/router"
	databasenani "backend/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var ServerContextInstance *ServiceContext

type ServiceContext struct {
	Config    *config.Config
	Server    *gin.Engine
	MySQLConn *gorm.DB
}

func NewServiceContext(path string) {
	cfg := config.MustLoadCfg(path, "YAML")
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false, // 禁止彩色打印
		},
	)
	gormConf := &gorm.Config{
		Logger: gormLogger,
	}
	dbCfg := cfg.DB
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database, dbCfg.Charset)
	err := databasenani.SetUpDB(gormConf, mysqlDsn, dbCfg.MaxIdleConns, dbCfg.MaxOpenConns, models.ModelList...)
	if err != nil {
		panic(err)
		return
	}
	ServerContextInstance = &ServiceContext{
		Config:    cfg,                  //
		MySQLConn: databasenani.GetDB(), // 集成MySQL查询对象
		Server:    router.SetUpRouter(),
	}
}

func GetSvcContext() *ServiceContext {
	return ServerContextInstance
}
