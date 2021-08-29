package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (m MovieRepositoryImpl) Save(db *gorm.DB, movie domain.Movie) domain.Movie {
	tx := db.Begin()

	if err := tx.Create(&movie).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()
	return movie
}

func (m MovieRepositoryImpl) Update(db *gorm.DB, movie domain.Movie) domain.Movie {
	panic("implement me")
}

func (m MovieRepositoryImpl) Delete(db *gorm.DB, movie domain.Movie) {
	panic("implement me")
}

func (m MovieRepositoryImpl) FindById(db *gorm.DB, id int) domain.Movie {
	panic("implement me")
}
