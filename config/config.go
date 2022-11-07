package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var loadedConfig *Config

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type TestConfig struct {
	Code       []uint   `json:"code"`
	Message    []string `json:"message"`
	Visibility string   `json:"visibility"`
}

type HTTPConfig struct {
	Port string `json:"port"`
}

type SocketConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	HTTP            HTTPConfig   `json:"http"`
	CacheTimeInMins int64        `json:"cache_time_in_mins"`
	AllowOrigins    []string     `json:"allow_origins"`
	AllowHeaders    []string     `json:"allow_headers"`
	Socket          SocketConfig `json:"socket"`
}

type DataBaseConfig struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
}

type HttpServerConnection struct {
	Domain string `json:"domain"`
	Path   string `json:"path"`
}

type Config struct {
	Server         ServerConfig   `json:"server"`
	Environment    string         `json:"env"`
	DataBaseConfig DataBaseConfig `json:"data_base_config"`
	JwtKey         string         `json:"jwt_key"`
	SecretKey      string         `json:"secret_key"`
}

func SetConfig(config *Config) {
	loadedConfig = config
}

func GetConfig() *Config {
	if loadedConfig == nil {
		config := &Config{}
		raw, err := ioutil.ReadFile("assets/config/conf.json")
		if err != nil {
			log.Fatalf("Error occured while reading config")
		}
		json.Unmarshal(raw, config)
		loadedConfig = config
	}
	return loadedConfig
}
