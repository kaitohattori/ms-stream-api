package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile    string
	APITimeout time.Duration

	AssetsDirPath      string
	AssetsM3u8FileName string
}

// Config is variable of ConfigList
var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("Failed to read file: ", err)
		os.Exit(1)
	}

	Config = ConfigList{
		LogFile:            cfg.Section("api").Key("log_file").String(),
		APITimeout:         time.Duration(cfg.Section("api").Key("api_timeout_sec").MustInt()) * time.Second,
		AssetsDirPath:      cfg.Section("file").Key("dir_path").String(),
		AssetsM3u8FileName: cfg.Section("file").Key("m3u8_filename").String(),
	}
}
