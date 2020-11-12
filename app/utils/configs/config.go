package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config = config
type Config struct {
	AppPort string
	Driver  string
	Auth    string
}

// Env = env
type Env struct {
	APISecret  string
	ServerPort string
}

// Load = load
func (c *Config) Load(driver string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	switch driver {
	case "mysql":
		dbHost := viper.GetString("MYSQL.DB_HOST")
		dbPort := viper.GetString("MYSQL.DB_PORT")
		dbUser := viper.GetString("MYSQL.DB_USER")
		dbPass := viper.GetString("MYSQL.DB_PASSWORD")
		dbName := viper.GetString("MYSQL.DB_NAME")

		auth := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

		c.AppPort = viper.GetString("APP.SERVER_PORT")
		c.Driver = viper.GetString("MYSQL.DB_DRIVER")
		c.Auth = auth

	case "postgres":
		// dbHost := viper.GetString("MYSQL.DB_HOST")
		// dbPort := viper.GetString("MYSQL.DB_PORT")
		// dbUser := viper.GetString("MYSQL.DB_USER")
		// dbPass := viper.GetString("MYSQL.DB_PASSWORD")
		// dbName := viper.GetString("MYSQL.DB_NAME")

		// auth := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

		// conn := Config{
		// 	AppPort: viper.GetString("APP.SERVER_PORT"),
		// 	Driver:  viper.GetString("MYSQL.DB_DRIVER"),
		// 	Auth:    auth,
		// }
	}

	return c, nil
}

// LoadApp = loadApp
func (e *Env) LoadApp() (*Env, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		return &Env{}, err
	}
	e.ServerPort = viper.GetString("APP.SERVER_PORT")
	e.APISecret = viper.GetString("APP.API_SECRET")
	return e, nil
}
