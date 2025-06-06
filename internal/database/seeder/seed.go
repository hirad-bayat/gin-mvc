package seeder

import (
	"fmt"
	Article "gin-demo/internal/modules/article/models"
	User "gin-demo/internal/modules/user/models"
	"gin-demo/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Seed() {
	db := database.Connection()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hashed password error")
		return
	}

	// Create a user
	user := User.User{
		Name:     "Random Name",
		Email:    "random@email.com",
		Role:     "customer",
		Password: string(hashedPassword),
	}
	result := db.Create(&user) // GORM sets user.ID after this
	if result.Error != nil {
		log.Fatalf("failed to create user: %v", result.Error)
	}
	log.Printf("user created successfully with email address: %s \n", user.Email)

	// Create articles linked to that user
	for i := 1; i <= 20; i++ {
		article := Article.Article{
			Title:   fmt.Sprintf("Random title %d", i),
			Content: fmt.Sprintf("Random content %d", i),
			UserID:  user.ID, // Use the actual ID
		}
		if err := db.Create(&article).Error; err != nil {
			log.Printf("failed to create article %d: %v", i, err)
		} else {
			log.Printf("article created successfully with title: %s \n", article.Title)
		}
	}
}
