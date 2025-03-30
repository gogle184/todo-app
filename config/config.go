package config

import (
	"log"

	"github.com/go-ini/ini"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").MustString("sqlite3"),
		DbName: cfg.Section("db").Key("name").MustString("webapp.sql"),
		LogFile: cfg.Section("web").Key("logfile").MustString("webapp.log"),
	}


}
