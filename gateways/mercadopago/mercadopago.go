package mercadopago

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApiUrl          string
	ApiVersion      string
	ApiClientID     string
	ApiClientSecret string
	ApiToken        string
	Charset         string
}

var conf *Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/paygo/config/")
	viper.AddConfigPath("$HOME/.paygo/config/")
	viper.AddConfigPath("$HOME/go/src/github/golangba/PayGo-Gateway/config/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetConfig() (*Config, error) {
	if conf != nil {
		return conf, nil
	}
	conf := new(Config)
	conf.ApiUrl = viper.GetString("mercadopago.config.apiUrl")
	conf.ApiVersion = viper.GetString("mercadopago.config.apiVersion")
	conf.ApiClientID = viper.GetString("mercadopago.config.apiClientID")
	conf.ApiClientSecret = viper.GetString("mercadopago.config.apiClientSecret")
	conf.ApiToken = viper.GetString("mercadopago.config.apiToken")
	conf.Charset = viper.GetString("mercadopago.config.charset")

	return conf, nil
}

func GetRoute(r string) (string, error) {
	routeName := fmt.Sprintf("mercadopago.routes.%s", r)
	route := viper.GetString(routeName)
	if len(route) == 0 {
		return route, fmt.Errorf("empty route")
	}
	return route, nil
}
