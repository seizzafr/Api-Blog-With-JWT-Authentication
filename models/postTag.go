package models
import (

)

type PostTag struct {
	PostID uint  `json:"post_id"`
	TagID uint  `json:"tag_id"` 
    Posts Posts  `gorm:"foreignKey:PostID"`
    Tags  Tags   `gorm:"foreignKey:TagID"`
}