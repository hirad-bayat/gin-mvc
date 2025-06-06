package bootstrap

import (
	"gin-demo/internal/database/seeder"
	"gin-demo/pkg/config"
	"gin-demo/pkg/database"
)

func Seed() {
	config.Set()
	database.Connect()
	seeder.Seed()
}
