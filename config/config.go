package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var Config *Cfg

type Cfg struct {
	Port       string `env:"PORT" envDefault:":8888"`
	DBHostname string `env:"DB_HOSTNAME"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBDatabase string `env:"DB_DATABASE"`
	ApiKey     string `env:"API_KEY"`
}

func Init() {

	Config = &Cfg{}
	err := env.Parse(Config)
	if err != nil {
		panic(err)
	}
	hlog.Infof("%#v", Config)
	InitDB()
}
