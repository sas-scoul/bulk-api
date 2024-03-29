package config

import "github.com/kelseyhightower/envconfig"

var Mysql MysqlConfigStruct

func init() {
	var mysql MysqlConfigStruct
	err := envconfig.Process("", &mysql)
	if err != nil {
		panic("Error loading env file: " + err.Error())
	}
	Mysql = mysql
}

type MysqlConfigStruct struct {
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Name     string `envconfig:"DB_NAME"`
}
