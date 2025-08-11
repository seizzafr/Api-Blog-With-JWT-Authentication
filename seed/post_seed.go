package seed

import (
	"api/config"
	"api/factory"
	"fmt"
)

func SeedPosts(count int) {
	db := config.DB

	// Ambil ID kategori dan user dari DB
	factory.LoadCategoryAndUserIDs()

	posts := factory.BatchPostFactory(count)

	for _, post := range posts {
		err := db.Create(&post).Error
		if err != nil {
			fmt.Printf("Gagal insert post: %v\n", err)
		} else {
			fmt.Printf("Post berhasil dibuat: %s\n", post.Title)
		}
	}
}
