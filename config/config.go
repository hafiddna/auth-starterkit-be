package config

import "github.com/spf13/viper"

type CfgStruct struct {
	App struct {
		Name        string `mapstructure:"name"`
		ServerName  string `mapstructure:"server_name"`
		Version     string `mapstructure:"version"`
		Environment string `mapstructure:"environment"`
		Server      struct {
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
		JWT struct {
			PublicKey            string `mapstructure:"public"`
			PrivateKey           string `mapstructure:"private"`
			RememberTokenPublic  string `mapstructure:"remember_token_public"`
			RememberTokenPrivate string `mapstructure:"remember_token_private"`
			GuestSessionSecret   string `mapstructure:"guest_session_secret"`
		} `mapstructure:"jwt"`
		Discord struct {
			Token string `mapstructure:"token"`
		} `mapstructure:"discord"`
		Secret struct {
			AuthKey          string `mapstructure:"auth_key"`
			RememberTokenKey string `mapstructure:"remember_token_key"`
		} `mapstructure:"secret"`
	} `mapstructure:"app"`
}

var Config CfgStruct

func GetConfig() (cfg CfgStruct, err error) {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.AddConfigPath(".")
	conf.SetConfigType("yaml")

	err = conf.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = conf.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
