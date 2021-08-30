package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct{}

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

func (m MovieRepositoryImpl) Update(db *gorm.DB, movie domain.Movie, slug string) domain.Movie {
	movieUpdated := domain.Movie{}
	db.First(&movieUpdated, "slug = ?", slug)
	movieUpdated.Title = movie.Title
	movieUpdated.Image = movie.Image
	movieUpdated.Slug = movie.Slug
	movieUpdated.Duration = movie.Duration
	movieUpdated.Description = movie.Description
	db.Save(&movieUpdated)

	return movieUpdated
}

func (m MovieRepositoryImpl) Delete(db *gorm.DB, slug string) {
	movie := domain.Movie{}
	db.Where("slug = ?", slug).Delete(&movie)
}

func (m MovieRepositoryImpl) FindBySlug(db *gorm.DB, slug string) domain.Movie {
	movie := domain.Movie{}
	db.First(&movie, "slug = ?", slug)

	return movie
}
