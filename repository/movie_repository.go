package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(db *gorm.DB, movie domain.Movie) domain.Movie
	Update(db *gorm.DB, movie domain.Movie) domain.Movie
	Delete(db *gorm.DB, movie domain.Movie)
	FindBySlug(db *gorm.DB, slug string) domain.Movie
}