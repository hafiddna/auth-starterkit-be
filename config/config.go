package config

import "github.com/spf13/viper"

type CfgStruct struct {
	App struct {
		Name         string `mapstructure:"name"`
		Version      string `mapstructure:"version"`
		Environtment string `mapstructure:"environment"`
		Server       struct {
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
			Cors string `mapstructure:"cors"`
		} `mapstructure:"server"`
		Redis struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Password string `mapstructure:"password"`
		} `mapstructure:"redis"`
		MongoDB struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Database string `mapstructure:"database"`
		} `mapstructure:"mongodb"`
		PostgreSQL struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Database string `mapstructure:"database"`
		} `mapstructure:"postgresql"`
		Minio struct {
			Host      string `mapstructure:"host"`
			Port      string `mapstructure:"port"`
			AccessKey string `mapstructure:"access_key"`
			SecretKey string `mapstructure:"secret_key"`
		} `mapstructure:"minio"`
		RabbitMQ struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"rabbitmq"`
		JWT struct {
			PublicKey  string `mapstructure:"public"`
			PrivateKey string `mapstructure:"private"`
		} `mapstructure:"jwt"`
	} `mapstructure:"app"`
}

type Config interface {
	GetConfig() CfgStruct
}

type config struct {
}

func NewConfig() Config {
	return &config{}
}

func (c *config) GetConfig() CfgStruct {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.AddConfigPath(".")
	conf.SetConfigType("yaml")

	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg CfgStruct
	err = conf.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
