package bootstrap

import (
	"gin-demo/internal/database/migration"
	"gin-demo/pkg/config"
	"gin-demo/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
