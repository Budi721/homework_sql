package repository

type MovieRepository interface {
	Save()
	Update()
	Delete()
	FindById()
	FindAll()
}