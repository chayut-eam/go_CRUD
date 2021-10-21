package app

import (
	"bytes"
	_ "embed"
	"strings"

	"crud/chrono"
	"crud/db"
	"crud/health"
	log "crud/logger"
	"crud/model"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var config []byte

func Bootstrap(appConfig *model.AppConfig) {
	// log
	log.Init(appConfig.AppInfo, appConfig.LoggingConfig)

	// chrono
	chrono.SetDefaultZone(chrono.Zone_Bangkok)

	// mongodb
	if err := db.Init(appConfig.MongoConfig); err != nil {
		log.Logger().Panicf("Fatal error connecting to database: %v", err)
	}

	// health check
	health.Init(appConfig.HealthCheckConfig)
}

func TearDown() {
	db.TearDown()
}

func LoadConfig() (*model.AppConfig, error) {
	// configuration type
	viper.SetConfigType("yaml")

	// environment variable
	viper.SetEnvPrefix("TOPVALUE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__", "-", "_"))
	viper.AutomaticEnv()

	// example:
	// TOPVALUE_UPSTREAM__BASE_URL -> upstream.base-url

	// read config
	if err := viper.ReadConfig(bytes.NewBuffer(config)); err != nil {
		return nil, err
	}

	// parse config into struct
	appConfig := &model.AppConfig{}
	if err := viper.Unmarshal(appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}
