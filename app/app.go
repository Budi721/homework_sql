package app

import (
	"github.com/Budi721/homework_sql/config"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DbConn *gorm.DB
}

func Init() *Application {
	db, err := InitDB()
	if err == nil {
		Migrate(db)
	}

	application := &Application{
		Config: config.Init(),
		DbConn: db,
	}

	return application
}
