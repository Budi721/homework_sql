package domain

type Movie struct {
	Id          uint
	Title       string `gorm:"type: VARCHAR(255) NOT NULL"`
	Slug        string `gorm:"type: VARCHAR(255) NOT NULL;unique"`
	Description string `gorm:"type:TEXT NOT NULL"`
	Duration    int    `gorm:"type: INT(5) NOT NULL"`
	Image       string `gorm:"type: VARCHAR(255) NOT NULL"`
}
