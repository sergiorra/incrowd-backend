package config

import (
	"encoding/json"
	"os"

	"github.com/labstack/gommon/log"
)

type Config struct {
	App                      App
	Database                 Database
	ArticlePoller            ArticlePoller
	Api                      Api
	Http                     Http
	HuddersfieldTownProvider Provider
}

type App struct {
	LogLevel                 string
	ShutdownTimeoutInSeconds int
}

type Database struct {
	User                  string
	Password              string
	Host                  string
	Port                  int
	SetupTimeoutInSeconds int
}

type ArticlePoller struct {
	ExecutionIntervalInMinutes int
	ExecutionTimeoutInSeconds  int
}

type Api struct {
	Host                       string
	Port                       int
	TimeoutInSeconds           int
	ReadHeaderTimeoutInSeconds int
}

type Http struct {
	MaxIdleConns        int
	MaxConnsPerHost     int
	MaxIdleConnsPerHost int
	TimeoutInSeconds    int
}

type Provider struct {
	Host   string
	Count  int
	TeamID string
}

func Read(filename string) (*Config, error) {
	var config Config

	log.Infof("Loading configuration from [%s]", filename)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
