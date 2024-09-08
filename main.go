package main

import (
	"apifoodweb/src/database"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Environment string

func init() {
	initLogrus()
	initViper()
}

func main() {
	logrus.Infoln("application start: this apps work on " + Environment + " environment.")

	logrus.Infoln("initiate database connection")
	if errInitConn := database.InitAllConnection(); errInitConn != nil {
		logrus.Error("error initiate database connction", errInitConn)
	}

	defer func() {
		logrus.Debug("close all database connection")
		database.CloseAllConnection()
		logrus.Error("application shutdown")
	}()
}

func initViper() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")

	if errConfig := viper.ReadInConfig(); errConfig != nil {
		logrus.Panicln("failed to read config:", errConfig)
	}

	Environment = "\033[31m" + viper.GetString("Environment") + "\033[0m"

}

func initLogrus() {

	logrus.SetOutput(os.Stdout)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableSorting:  false,
		PadLevelText:    true,
	})
}
