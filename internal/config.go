package config

import (
	"apifoodweb/pkg/util"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string `json:"DB_HOST"`
	User     string `json:"DB_USERNAME"`
	Password string `json:"DB_PASSWORD"`
	Database string `json:"DB_NAME"`
	Port     string `json:"DB_PORT"`
	SslMode  string `json:"SSL_MODE"`
}

type AppConfig struct {
	ApiConfig
	DbConfig
}

// Turn struct postgress database config into connection string.
func (c *AppConfig) connString() string {
	util.SystemLog("Init Connection", "transform struct to connection string", nil).Debug()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, util.DecryptCamellia(c.Password), c.Database, c.Port)
}

// Function to get specific postgress database config,insert viper direction as input parameter.
func (c *AppConfig) readFileConfig(viperDirection string) {

	var configStruct DbConfig

	util.SystemLog("Init Connection", "Initiate connection. try get connection string use viper.", nil).Debug()

	configJson := viper.GetStringMap(viperDirection)

	bytedConfig, errMarshal := json.Marshal(configJson)
	if errMarshal != nil {
		util.SystemLog("Init Connection", "failed marshal config json", errMarshal).Error()
		return
	}

	if errUnmarshal := json.Unmarshal(bytedConfig, &configStruct); errUnmarshal != nil {
		util.SystemLog("Init Connection", "failed unmarshal config json", errUnmarshal).Error()
		return
	}

	configStruct.Password = util.DecryptCamellia(configStruct.Password)

	util.SystemLog("Init Connection", "Success get config "+configStruct.Database+" database", nil).Debug()
	c.DbConfig = configStruct
	c.ApiConfig.ApiPort = viper.GetString("SERVER_PORT")

	return
}

func NewConfig(viperDirection string) AppConfig {
	configuration := AppConfig{}
	configuration.readFileConfig(viperDirection)
	return configuration
}
