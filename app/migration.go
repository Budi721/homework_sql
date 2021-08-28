package app

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(domain.Movie{})
	if err != nil {
		return
	}
}