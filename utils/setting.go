package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassport string
	DbName     string

	COS_BUCKET_NAME string
	COS_REGION      string
	COS_APP_ID      string
	COS_SECRET_ID   string
	COS_SECRET_KEY  string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadDatabase(cfg)
	LoadCos(cfg)
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
	JwtKey = sec.Key("JwtKey").MustString("progress_key")
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

func LoadCos(cfg *ini.File) {
	sec := cfg.Section("cos")
	COS_BUCKET_NAME = sec.Key("COS_BUCKET_NAME").MustString("progress-management")
	COS_REGION = sec.Key("COS_REGION").MustString("ap-chengdu")
	COS_APP_ID = sec.Key("COS_APP_ID").MustString("1304266993")
	COS_SECRET_ID = sec.Key("COS_SECRET_ID").String()
	COS_SECRET_KEY = sec.Key("COS_SECRET_KEY").String()
}
