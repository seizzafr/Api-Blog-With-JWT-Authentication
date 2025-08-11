package main

import (
    "api/config"
    "api/models"
    "api/routes"
    "log"
    // "api/seed"

	"github.com/joho/godotenv"
)


func main() {
  err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}
    config.ConnectDB()
    config.DB.AutoMigrate(
		&models.Category{},
		&models.Posts{},
    &models.Users{},
    &models.Tags{},
    &models.PostTag{},
		)

    //  seed.SeedUsers(10);
    //  seed.SeedCategories(10)
    //  seed.SeedPosts(10)
    r := routes.SetupRoutes()
    r.Run(":8080")
}
