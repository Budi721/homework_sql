package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(db *gorm.DB, movie domain.Movie) domain.Movie
	Update(db *gorm.DB, movie domain.Movie, slug string) domain.Movie
	Delete(db *gorm.DB, slug string)
	FindBySlug(db *gorm.DB, slug string) domain.Movie
}