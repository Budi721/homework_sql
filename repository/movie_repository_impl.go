package repository

import (
	"github.com/Budi721/homework_sql/model/domain"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct{}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (m MovieRepositoryImpl) Save(db *gorm.DB, movie domain.Movie) (domain.Movie, error) {
	tx := db.Begin()

	if err := tx.Create(&movie).Error; err != nil {
		tx.Rollback()
		return domain.Movie{}, err
	}

	tx.Commit()
	return movie, nil
}

func (m MovieRepositoryImpl) Update(db *gorm.DB, movie domain.Movie, slug string) (domain.Movie, error) {
	movieUpdated := domain.Movie{}
	db.First(&movieUpdated, "slug = ?", slug)
	movieUpdated.Title = movie.Title
	movieUpdated.Image = movie.Image
	movieUpdated.Slug = movie.Slug
	movieUpdated.Duration = movie.Duration
	movieUpdated.Description = movie.Description
	if err := db.Save(&movieUpdated).Error; err != nil {
		return domain.Movie{}, err
	}

	return movieUpdated, nil
}

func (m MovieRepositoryImpl) Delete(db *gorm.DB, slug string) error {
	movie := domain.Movie{}
	err := db.First(&movie, "slug = ?", slug).Error
	if err != nil {
		return err
	}

	db.Where("slug = ?", slug).Delete(&movie)
	return nil
}

func (m MovieRepositoryImpl) FindBySlug(db *gorm.DB, slug string) (domain.Movie, error) {
	movie := domain.Movie{}
	err := db.First(&movie, "slug = ?", slug).Error
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}
