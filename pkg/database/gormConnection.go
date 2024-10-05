package database

import (
	"apifoodweb/pkg/util"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type poolGorm struct {
	gormConn       *gorm.DB
	database       string
	viperDirection string
}

var (
	GormApifoodapp *poolGorm
)

func GetConnectionGormApiFoodApp() (*gorm.DB, error) {
	return GormApifoodapp.checkConnectionGorm()
}

func InitAllGormConnection() error {

	GormApifoodapp = new(poolGorm)

	errOpen := createConnectionGorm(GormApifoodapp, "Database.Postgres.apifoodapp")
	if errOpen != nil {
		return errOpen
	}

	return nil
}

func CloseAllGormConnection() {

	apiFoodAppConn, _ := GormApifoodapp.gormConn.DB()
	apiFoodAppConn.Close()

}

// Call this function to open database connection using gorm and asign connection to connection pool, insert viper direction as input parameter.
func createConnectionGorm(dbConn *poolGorm, viperDirection string) error {

	config := getPostgresConfig(viperDirection)

	connString := config.connString()

	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)

	dbConn.gormConn = db
	dbConn.database = config.Database
	dbConn.viperDirection = viperDirection

	util.SystemLog("Init Connection", "Success open and store connection to pool: database "+config.Database+" opened", nil).Debug()

	return nil
}

// Function for check connection gorm, retry to reopen connection while the connection closed.
func (db *poolGorm) checkConnectionGorm() (*gorm.DB, error) {

	if db.gormConn == nil {
		util.SystemLog("Get Connection", "connection "+db.database+" not found, reopen connection", nil).Debug()
		createConnectionGorm(db, db.viperDirection)
	}

	for index := 0; index < 5; index++ {

		var errPing error
		conn, _ := db.gormConn.DB()

		if errPing = conn.Ping(); errPing != nil {

			util.SystemLog("Get Connection", "connection "+ConnApiFoodApp.database+" not responded", errPing).Debug()
			createConnectionGorm(db, db.viperDirection)

			if errPing == nil {
				util.SystemLog("Get Connection", "connection "+ConnApiFoodApp.database+" responded", nil).Debug()
				break
			}

			if errPing != nil && index == 4 {
				util.SystemLog("Get Connection", "failed to connect "+ConnApiFoodApp.database+" database", errPing).Debug()
				return nil, errPing
			}

			time.Sleep(1 * time.Minute)
		}
		break

	}
	return db.gormConn, nil
}
