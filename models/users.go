package models

import (
	"time"
)

type Users struct {
	ID  uint `gorm:"primary_key" json:"id"`    
	Username  string `json:"username"`
	Email  string `json:"email" gorm:"type:varchar(225);unique"`
	Password string `json:"password"`
	CreatedAt  time.Time `json:"created_at"` 
	UpdatedAt  time.Time `json:"updated_at"`  
}