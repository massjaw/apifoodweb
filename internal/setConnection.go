package config

import (
	"apifoodweb/src/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"dbname"`
	Port     int    `json:"port"`
}

type poolPgsql struct {
	dbPgsql        *sql.DB
	database       string
	viperDirection string
}

var (
	ConnApiFoodApp *poolPgsql
)

// ======================== Get connection that already open in InitAllConnection function ===========================================

// Function to get connection apifoodapp database from connection pool
func GetConnectionApiFoodApp() (*sql.DB, error) {

	util.SystemLog("Get Connection", "Getting connection for database apifoodapp", nil).Info()
	return ConnApiFoodApp.checkConnectionPostgres()
}

// ========================= Compiller for open all connection and store connection inside the pool =====================================

func InitAllConnection() error {
	var errCreateConnection error

	ConnApiFoodApp = new(poolPgsql)

	if errCreateConnection = createConnectionPostgres(ConnApiFoodApp, "Database.Postgres.apifoodapp"); errCreateConnection != nil {
		return errCreateConnection
	}

	return nil
}

func CloseAllConnection() {
	ConnApiFoodApp.dbPgsql.Close()
}

// ======================================= Connection builder to open connection ===================================================

// Call this function to open database connection and asign connection to connection pool, insert viper direction as input parameter.
func createConnectionPostgres(dbConn *poolPgsql, viperDirection string) error {

	config := getPostgresConfig(viperDirection)

	connString := config.connString()

	var db *sql.DB
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	dbConn.dbPgsql = db
	dbConn.database = config.Database
	dbConn.viperDirection = viperDirection

	util.SystemLog("Init Connection", "Success open and store connection to pool: database "+config.Database+" opened", nil).Debug()

	return nil
}

// =========================================== local function to get connection string from config.json =========================================

// Turn struct postgress database config into connection string.
func (c *PostgresConfig) connString() string {
	util.SystemLog("Init Connection", "transform struct to connection string", nil).Debug()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Host, c.Username, util.DecryptCamellia(c.Password), c.Database, c.Port)
}

// Function to get specific postgress database config,insert viper direction as input parameter.
func getPostgresConfig(viperDirection string) *PostgresConfig {

	var configStruct PostgresConfig

	util.SystemLog("Init Connection", "Initiate connection. try get connection string use viper.", nil).Debug()

	configJson := viper.GetStringMap(viperDirection)

	bytedConfig, errMarshal := json.Marshal(configJson)
	if errMarshal != nil {
		util.SystemLog("Init Connection", "failed marshal config json", errMarshal).Error()
		return nil
	}

	if errUnmarshal := json.Unmarshal(bytedConfig, &configStruct); errUnmarshal != nil {
		util.SystemLog("Init Connection", "failed unmarshal config json", errUnmarshal).Error()
		return nil
	}

	util.SystemLog("Init Connection", "Success get config "+configStruct.Database+" database", nil).Debug()
	return &configStruct
}

// Function for check connection, retry to reopen connection while the connection closed.
func (db *poolPgsql) checkConnectionPostgres() (*sql.DB, error) {

	if db.dbPgsql == nil {
		util.SystemLog("Get Connection", "connection "+db.database+" not found, reopen connection", nil).Debug()
		createConnectionPostgres(db, db.viperDirection)
	}

	for index := 0; index < 5; index++ {

		var errPing error

		if errPing = db.dbPgsql.Ping(); errPing != nil {

			time.Sleep(1 * time.Minute)

			util.SystemLog("Get Connection", "connection "+ConnApiFoodApp.database+" not responded", errPing).Debug()
			createConnectionPostgres(db, db.viperDirection)

			if errPing == nil {
				util.SystemLog("Get Connection", "connection "+ConnApiFoodApp.database+" responded", nil).Debug()
				break
			}

			if errPing != nil && index == 4 {
				util.SystemLog("Get Connection", "failed to connect "+ConnApiFoodApp.database+" database", errPing).Debug()
				return nil, errPing
			}
		}

	}
	return db.dbPgsql, nil
}
