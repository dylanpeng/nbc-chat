package config

import (
	"github.com/BurntSushi/toml"
	oConf "github.com/dylanpeng/nbc-chat/common/config"
	"github.com/dylanpeng/nbc-chat/lib/chatgpt"
)

var conf *Config

type Config struct {
	*oConf.Config
	ChatGPT *chatgpt.Config `toml:"chatgpt" json:"chatgpt"`
}

func Init(file string) error {
	conf = &Config{
		Config: oConf.Default(),
	}

	if _, err := toml.DecodeFile(file, conf); err != nil {
		return err
	}

	if err := conf.Config.Init(); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return conf
}
