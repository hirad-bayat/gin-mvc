package config

import "gin-demo/config"

func Get() config.Config {
	return configurations
}
