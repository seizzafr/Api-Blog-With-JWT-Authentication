package seed

import (
	"fmt"

	"api/config"
	"api/factory"
)

func SeedUsers(count int) {
	db := config.DB
	users := factory.BatchUserFactory(count)

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			fmt.Printf("❌ Gagal insert user: %v\n", err)
		} else {
			fmt.Printf("✅ User %s berhasil disimpan\n", user.Email)
		}
	}
}
