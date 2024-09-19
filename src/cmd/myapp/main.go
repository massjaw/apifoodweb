package main

import (
	"apifoodweb/config"
	database "apifoodweb/config"
	"apifoodweb/src/models"
	"apifoodweb/src/util"
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

	db, errDb := config.GetConnectionGormApiFoodApp()
	if errDb != nil {
		logrus.Error("failed to get gorm connection", errDb)
	}

	// Drop existing tables if they exist
	db.Migrator().DropTable(&models.UserDetail{})
	db.Migrator().DropTable(&models.Users{})

	// AutoMigrate will create the tables if they don't exist
	err := db.AutoMigrate(&models.Users{}, &models.UserDetail{})
	if err != nil {
		logrus.Error("failed to migrate database: %v", err)
	}

	// Sample operations
	// Create a user with detail
	user := models.Users{
		Username:       "massjaw",
		Email:          "ibrahimhisyam@gmail.com",
		HashedPassword: util.HashPassword("IniPassword"),
	}
	db.Create(&user)

	userDetail := models.UserDetail{
		UserID:         user.ID,
		FirstName:      "Ibrahim",
		MiddleName:     "Muhammad",
		LastName:       "Hisyam",
		Address:        "Jalan jalan aja gih sonoo",
		PhoneNumber:    "087724400122",
		ProfilePicture: "ibrahim.jpg",
	}
	db.Create(&userDetail)

	// Retrieve the user with user detail
	var getUser models.Users
	var getUserDetail models.UserDetail
	db.First(&getUser)
	db.First(&getUserDetail)

	logrus.Info("User: ", getUser)
	logrus.Info("UserDetail: ", getUserDetail)

	defer func() {
		logrus.Debug("close all database connection")
		database.CloseAllGormConnection()
		logrus.Debug("application shutdown")
	}()
}

func initViper() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("backend/config/")

	if errConfig := viper.ReadInConfig(); errConfig != nil {
		logrus.Panicln("failed to read config:", errConfig)
	}

	Environment = "\033[31m" + viper.GetString("Environment") + "\033[0m"

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
