package seed

import (
	"fmt"
	"api/factory"
	"api/config"
)

func SeedCategories(count int) {
	db := config.DB

	categories := factory.BatchCategoryFactory(count)
	for _, cat := range categories {
		err := db.Create(&cat).Error
		if err != nil {
			fmt.Printf("Gagal insert kategori: %v\n", err)
		} else {
			fmt.Printf("Kategori berhasil dibuat: %s\n", cat.Name)
		}
	}
}
