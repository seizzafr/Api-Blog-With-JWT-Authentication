package models

import (
	"time"
	"api/utils"
	"gorm.io/gorm"
)
type Posts struct {
	ID         uint      `gorm:"primary_key" json:"id"`        // Primary key, biasanya uint
	CategoryID uint      `json:"category_id"`  // Foreign key ke kategori
	Title      string    `json:"title"`        // Judul artikel
	Content    string    `json:"content"`      // Isi artikel
	Thumbnail  string    `json:"thumbnail"`    // URL atau path gambar thumbnail
	UserID     uint      `json:"user_id"`      // Foreign key ke user (penulis)
	CreatedAt  time.Time `json:"created_at"`   // Waktu dibuat
	UpdatedAt  time.Time `json:"updated_at"`   // Waktu diperbarui
    Slug string `gorm:"unique;not null"` // <- kolom baru
	Category   Category  `gorm:"foreignKey:CategoryID"`
    Users      Users     `gorm:"foreignKey:UserID"`
}

func (p *Posts) BeforeSave(tx *gorm.DB) (err error) {
    p.Slug = utils.GenerateUniqueSlug(tx, p.Title, &Posts{}, "slug")
    return
}


