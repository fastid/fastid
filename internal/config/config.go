package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	Config struct {
		Debug    bool `yaml:"debug" env:"DEBUG"`
		HTTP     `yaml:"http"`
		DATABASE `yaml:"database"`
		LOGGER   `yaml:"logger"`
	}

	LOGGER struct {
		Level string `env-required:"true" yaml:"level" env:"LOGGER_LEVEL"`
	}

	HTTP struct {
		Listen string `env-required:"true" yaml:"listen" env:"HTTP_LISTEN"`
	}

	DATABASE struct {
		DriverName      string `env-default:"driver_name" yaml:"driver_name" env:"DATABASE_DRIVER_NAME"`
		DBName          string `env-default:"fastid" yaml:"dbname" env:"DATABASE_DBNAME"`
		User            string `env-default:"user" yaml:"user" env:"DATABASE_USER"`
		Password        string `env-default:"password" yaml:"password" env:"DATABASE_PASSWORD"`
		Host            string `env-default:"localhost" yaml:"host" env:"DATABASE_HOST"`
		Port            string `env-default:"5432" yaml:"port" env:"DATABASE_PORT"`
		SslMode         string `env-default:"disable" yaml:"sslmode" env:"DATABASE_SSLMODE"`
		ApplicationName string `env-default:"FastID" yaml:"application_name" env:"DATABASE_APPLICATION_NAME"`
		Scheme          string `env-default:"public" yaml:"scheme" env:"DATABASE_SCHEME"`
		ConnectTimeout  string `env-default:"10" yaml:"connect_timeout" env:"DATABASE_CONNECTION_TIMEOUT"`
		MaxOpenConns    int    `env-default:"20" yaml:"max_open_conns" env:"DATABASE_MAX_OPEN_CONNS"`
		MaxIdleConns    int    `env-default:"5" yaml:"max_idle_conns" env:"DATABASE_MAX_IDLE_CONNS"`
		ConnMaxLifetime int    `env-default:"1800" yaml:"conn_max_lifetime" env:"DATABASE_MAX_LIFETIME"`
		ConnMaxIdleTime int    `env-default:"1800" yaml:"conn_max_idletime" env:"DATABASE_MAX_IDLETIME"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	fileInfo, _ := os.Stat("configs/fastid.yml")
	if fileInfo != nil {
		if err := cleanenv.ReadConfig("configs/fastid.yml", cfg); err != nil {
			return nil, err
		}
	}

	fileInfo, _ = os.Stat("fastid.yml")
	if fileInfo != nil {
		if err := cleanenv.ReadConfig("fastid.yml", cfg); err != nil {
			return nil, err
		}
	}

	fileInfo, _ = os.Stat("../../configs/fastid.yml")
	if fileInfo != nil {
		if err := cleanenv.ReadConfig("../../configs/fastid.yml", cfg); err != nil {
			return nil, err
		}
	}

	fileInfo, _ = os.Stat(".env")
	if fileInfo != nil {
		_ = cleanenv.ReadConfig(".env", cfg)
	}

	fileInfo, _ = os.Stat("../../.env")
	if fileInfo != nil {
		_ = cleanenv.ReadConfig("../../.env", cfg)
	}

	_ = cleanenv.ReadEnv(cfg)
	return cfg, nil
}
