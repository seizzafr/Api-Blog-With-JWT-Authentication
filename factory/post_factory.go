package factory

import (
	"api/config"
	"api/models"
	"github.com/bxcodec/faker/v4"
	"log"
	"math/rand"
	"time"
)

var categoryIDs []uint
var userIDs []uint

// Load ID dari tabel categories dan users
func LoadCategoryAndUserIDs() {
	var categories []models.Category
	var users []models.Users

	// Ambil semua ID kategori dari database
	if err := config.DB.Find(&categories).Error; err != nil {
		log.Fatalf("Gagal load categories: %v", err)
	}

	for _, cat := range categories {
		categoryIDs = append(categoryIDs, cat.ID)
	}

	// Ambil semua ID user dari database
	if err := config.DB.Find(&users).Error; err != nil {
		log.Fatalf("Gagal load users: %v", err)
	}

	for _, u := range users {
		userIDs = append(userIDs, u.ID)
	}

	// Jika kosong, munculkan error biar nggak insert kosong
	if len(categoryIDs) == 0 || len(userIDs) == 0 {
		log.Fatal("Data category atau user kosong. Seed kategori dan user terlebih dahulu.")
	}
}

// Generate 1 post
func NewPostFactory() models.Posts {
	return models.Posts{
		Title:      faker.Sentence(),
		Content:    faker.Paragraph(),
		Thumbnail:  faker.URL(),
		CategoryID: categoryIDs[rand.Intn(len(categoryIDs))],
		UserID:     userIDs[rand.Intn(len(userIDs))],
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// Generate banyak post
func BatchPostFactory(count int) []models.Posts {
	posts := make([]models.Posts, count)
	for i := 0; i < count; i++ {
		posts[i] = NewPostFactory()
	}
	return posts
}
