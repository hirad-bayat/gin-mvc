package migration

import (
	"fmt"
	Article "gin-demo/internal/modules/article/models"
	User "gin-demo/internal/modules/user/models"
	"gin-demo/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&User.User{}, &Article.Article{})
	if err != nil {
		log.Fatal("can't migrate")
		return
	}
	fmt.Println("migration done successfully...")
}
