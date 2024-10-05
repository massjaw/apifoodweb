package main

import (
	"apifoodweb/api/server"
	"apifoodweb/internal/database"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Environment string

func init() {
	initLogrus()
	initViper()
	initConnection()
}

func main() {
	logrus.Debug("application start: this apps work on " + Environment + " environment.")

	server.InitApplicationServer().Run()
}

func initViper() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("internal/")

	if errConfig := viper.ReadInConfig(); errConfig != nil {
		logrus.Panicln("failed to read config:", errConfig)
	}

	Environment = "\033[31m" + viper.GetString("ENVIRONMENT") + "\033[0m"

}

func initLogrus() {

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableSorting:  false,
		PadLevelText:    true,
	})
}

func initConnection() {
	logrus.Debug("initiate database connection")
	if errInitConn := database.InitAllGormConnection(); errInitConn != nil {
		logrus.Panic("error initiate database connction", errInitConn)
	}
}
