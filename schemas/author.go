package schemas

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Books []Book `gorm:"foreignKey:AuthorId" json:"books"`
}