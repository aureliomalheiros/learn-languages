package configs

import (
	"github.com/spf13/viper"
	"github.com/go-chi/jwtauth"
)

var cfg *conf

type conf struct {
	DBDRIVER 		string `mapstructure:"DB_DRIVER"`
	DBHOST 			string `mapstructure:"DB_HOST"`
	DBPORT 			string `mapstructure:"DB_PORT"`
	DBUSER 			string `mapstructure:"DB_USER"`
	DBPASS 			string `mapstructure:"DB_PASS"`
	DBNAME 			string `mapstructure:"DB_NAME"`
	WEBSERVERPORT 	string `mapstructure:"WEBSERVER_PORT"`
	JWTSecret 		string `mapstructure:JWT_SECRET`
	JwTExperesIn 	int    `mapstructure:JWT_EXPIRESIN`
	TokenAuth 		*jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SecConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}