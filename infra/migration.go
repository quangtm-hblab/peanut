package infra

import (
	"peanut/domain"

	"gorm.io/gorm"
)

func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&domain.User{}, &domain.Book{}, &domain.Content{})
	return
}
