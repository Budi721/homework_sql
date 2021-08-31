package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(db *gorm.DB, movie domain.Movie) (domain.Movie, error)
	Update(db *gorm.DB, movie domain.Movie, slug string) (domain.Movie, error)
	Delete(db *gorm.DB, slug string) error
	FindBySlug(db *gorm.DB, slug string) (domain.Movie, error)
}