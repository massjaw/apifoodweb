package database

import (
	"apifoodweb/pkg/util"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string `json:"DB_HOST"`
	Username string `json:"DB_USERNAME"`
	Password string `json:"DB_PASSWORD"`
	Database string `json:"DB_NAME"`
	Port     int    `json:"DB_PORT"`
}
type GormConnection struct {
	gormConn       *gorm.DB
	database       string
	viperDirection string
}

var (
	GormApifoodapp *GormConnection
)

func GetConnectionGormApiFoodApp() (*gorm.DB, error) {
	return GormApifoodapp.checkConnectionGorm()
}

func InitAllGormConnection() error {

	GormApifoodapp = new(GormConnection)

	errOpen := GormApifoodapp.createConnectionGorm("DATABASE.APIFOODAPP")
	if errOpen != nil {
		return errOpen
	}

	return nil
}

func CloseAllGormConnection() {

	apiFoodAppConn, _ := GormApifoodapp.gormConn.DB()
	apiFoodAppConn.Close()
	return
}

// Call this function to open database connection using gorm and asign connection to connection pool, insert viper direction as input parameter.
func (g *GormConnection) createConnectionGorm(viperDirection string) error {

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

	g.gormConn = db
	g.database = config.Database
	g.viperDirection = viperDirection

	util.SystemLog("Init Connection", "Success open and store connection to pool: database "+config.Database+" opened", nil).Debug()

	return nil
}

// Function for check connection gorm, retry to reopen connection while the connection closed.
func (g *GormConnection) checkConnectionGorm() (*gorm.DB, error) {

	if g.gormConn == nil {
		util.SystemLog("Get Connection", "connection "+g.database+" not found, reopen connection", nil).Debug()
		g.createConnectionGorm(g.viperDirection)
	}

	for index := 0; index < 5; index++ {

		var errPing error
		conn, _ := g.gormConn.DB()

		if errPing = conn.Ping(); errPing != nil {

			util.SystemLog("Get Connection", "connection "+g.database+" not responded", errPing).Debug()
			g.createConnectionGorm(g.viperDirection)

			if errPing == nil {
				util.SystemLog("Get Connection", "connection "+g.database+" responded", nil).Debug()
				break
			}

			if errPing != nil && index == 4 {
				util.SystemLog("Get Connection", "failed to connect "+g.database+" database", errPing).Debug()
				return nil, errPing
			}

			time.Sleep(1 * time.Minute)
		}
		break

	}
	return g.gormConn, nil
}

// Turn struct postgress database config into connection string.
func (c *PostgresConfig) connString() string {
	util.SystemLog("Init Connection", "transform struct to connection string", nil).Debug()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Host, c.Username, util.DecryptCamellia(c.Password), c.Database, c.Port)
}
