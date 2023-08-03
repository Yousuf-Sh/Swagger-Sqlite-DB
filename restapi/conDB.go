package restapi

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define the User model for GORM
type dbUser struct {
	gorm.Model
	Name  *string `gorm:"column:name"`
	Email *string `gorm:"column:email"`
}

func (dbUser) TableName() string {
	return "db_users"
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../userDB.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the User model to create the "users" table if it doesn't exist
	err = db.AutoMigrate(&dbUser{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
