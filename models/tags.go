package models

import (
"api/utils"
"gorm.io/gorm"
)

type Tags struct {
	ID  uint `gorm:"primary_key" json:"id"`
	Name string    `json:"name"`
	Slug string `gorm:"unique;not null"` 
}

func (p *Tags) BeforeSave(tx *gorm.DB) (err error) {
    p.Slug = utils.GenerateUniqueSlug(tx, p.Name, &Tags{}, "slug")
    return
}
