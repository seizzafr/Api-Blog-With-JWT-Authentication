package factory

import (
	"api/models"
	"time"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pw string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "" // kalau error, kosongin aja
	}
	return string(hashed)
}

// Buat satu user dummy
func NewUserFactory() models.Users {
	password := "password123" // password default dummy
	hashedPassword := hashPassword(password)

	return models.Users{
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Buat banyak user dummy
func BatchUserFactory(count int) []models.Users {
	users := make([]models.Users, count)
	for i := 0; i < count; i++ {
		users[i] = NewUserFactory()
	}
	return users
}
