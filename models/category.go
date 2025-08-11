package models
import ( 
	   "api/utils"
       "gorm.io/gorm"
)
type Category struct {
	ID   uint    `gorm:"primary_key" json:"categoryid"`
	Name string `json:"name"`
	Slug string `gorm:"unique;not null"` 
}

func (p *Category) BeforeSave(tx *gorm.DB) (err error) {
    p.Slug = utils.GenerateUniqueSlug(tx, p.Name, &Category{}, "slug")
    return
}
