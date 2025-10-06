package models

import "gorm.io/gorm"

type Todos struct {
	ID        uint    `gorm: "primary key;autoIncrement" json:"id"`
	Item      *string `json: "item"`
	Completed bool    `gorm:"default:false" json:"completed"`
}

func MigrateTodos(db *gorm.DB) error {
	err := db.AutoMigrate(&Todos{})
	return err
}
