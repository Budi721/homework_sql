package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(db *gorm.DB, movie domain.Movie) domain.Movie
	Update(db *gorm.DB, movie domain.Movie) domain.Movie
	Delete(db *gorm.DB, movie domain.Movie)
	FindById(db *gorm.DB, id int) domain.Movie
}