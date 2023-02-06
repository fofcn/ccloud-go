package config

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/goinggo/mapstructure"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type LoggingConfig struct {
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	LocalTime  bool   `yaml:"localTime"`
	Compress   bool   `yaml:"compress"`
}

type StoreConfig struct {
	ImagePath string `yaml:"image"`
	VideoPath string `yaml:"video"`
}

type DataSourceConfig struct {
	DriverName string       `yaml:"driverName"`
	Config     SqliteConfig `yaml:"config"`
}

type SqliteConfig struct {
	DbPath     string `yaml:"dbPath"`
	InitSqlDir string `yaml:"initSqlDir"`
}

type AccountConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ConfigLoader struct {
	ServerConfig     ServerConfig
	LoggingConfig    LoggingConfig
	StoreConfig      StoreConfig
	DataSourceConfig DataSourceConfig
	AccountConfig    AccountConfig
}

var configLoader *ConfigLoader
var once sync.Once

func GetInstance() *ConfigLoader {
	once.Do(func() {
		configLoader = &ConfigLoader{}
	})

	return configLoader
}

func (loader *ConfigLoader) LoadConfig() error {
	buf, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		log.Fatalf("Read env setup yaml file error, %v", err)
		return err
	}

	var config map[string]interface{}
	err = yaml.Unmarshal([]byte(buf), &config)
	if err != nil {
		log.Fatalf("Unmarshal yaml file error, %v", err)
		return err
	}

	err = loader.loadServerConfig(config)
	if err != nil {
		return err
	}

	err = loader.loadLoggingConfig(config)
	if err != nil {
		return err
	}

	err = loader.loadStoreConfig(config)
	if err != nil {
		return err
	}

	err = loader.loadDataSourceConfig(config)
	if err != nil {
		return err
	}

	err = loader.loadAccountConfig(config)
	if err != nil {
		return err
	}

	return err
}

func (loader *ConfigLoader) loadServerConfig(config map[string]interface{}) error {
	return mapstructure.Decode(config["server"], &loader.ServerConfig)
}

func (loader *ConfigLoader) loadLoggingConfig(config map[string]interface{}) error {
	return mapstructure.Decode(config["logging"], &loader.LoggingConfig)
}

func (loader *ConfigLoader) loadStoreConfig(config map[string]interface{}) error {
	return mapstructure.Decode(config["store"], &loader.StoreConfig)
}

func (loader *ConfigLoader) loadDataSourceConfig(config map[string]interface{}) error {
	return mapstructure.Decode(config["datasource"], &loader.DataSourceConfig)
}

func (loader *ConfigLoader) loadAccountConfig(config map[string]interface{}) error {
	return mapstructure.Decode(config["account"], &loader.AccountConfig)
}
