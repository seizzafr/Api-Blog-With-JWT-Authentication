package factory

import (
	"fmt"
	"math/rand"
	"api/models"
)

var categoryNames = []string{"Technology", "Health", "Education", "Finance", "Lifestyle"}

func NewCategoryFactory() models.Category {
	name := categoryNames[rand.Intn(len(categoryNames))]
	return models.Category{
		Name: name,
		Slug: fmt.Sprintf("%s-%d", name, rand.Intn(1000)),
	}
}

func BatchCategoryFactory(count int) []models.Category {
	categories := make([]models.Category, count)
	for i := 0; i < count; i++ {
		categories[i] = NewCategoryFactory()
	}
	return categories
}
