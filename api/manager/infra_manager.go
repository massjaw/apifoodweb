package manager

import (
	config "apifoodweb/internal"
	"apifoodweb/pkg/util"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	DbConn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg config.AppConfig
}

func NewInfraManager(cfg config.AppConfig) InfraManager {
	infra := infraManager{
		cfg: cfg,
	}
	infra.initDB()
	return &infra
}

func (i *infraManager) DbConn() *gorm.DB {
	return i.db
}

func (i *infraManager) initDB() {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Database, i.cfg.SslMode)
	gormDb, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			db, _ := gormDb.DB()
			util.SystemLog("Database Connection", "database failed to connected", nil).Fatal()
			db.Close()
		}
	}()

	i.db = gormDb
}
