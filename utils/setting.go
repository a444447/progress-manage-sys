package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassport string
	DbName     string
)

func init() {
	cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadDatabase(cfg)
}

func LoadServer(cfg *ini.File) {
	sec := cfg.Section("server")
	AppMode = sec.Key("AppMode").MustString("debug")
	HttpPort = sec.Key("HttpMode").Validate(func(in string) string {
		if len(in) == 0 {
			return ":3001"
		}
		return in
	})
}

func LoadDatabase(cfg *ini.File) {
	sec := cfg.Section("database")
	Db = sec.Key("Db").MustString("mysql")
	DbHost = sec.Key("DbHost").Validate(func(in string) string {
		if len(in) == 0 {
			return "localhost"
		}
		return in
	})
	DbPort = sec.Key("DbPort").MustString("3306")
	DbUser = sec.Key("DbUser").MustString("root")
	DbPassport = sec.Key("DbPassport").MustString("root")
	DbName = sec.Key("DbName").MustString("user")
}
