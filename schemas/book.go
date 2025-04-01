package schemas

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	AuthorId uint
	Title 	string         `gorm:"type:varchar(100);not null" json:"title"`
}