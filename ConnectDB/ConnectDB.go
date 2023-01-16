package ConnectDB

import (
	"github.com/gym/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDBLocation() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("././database/Location.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(models.Location{}); err != nil {
		return nil, err
	}
	return db, nil
}
func GetDBSchool() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("././database/School.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(models.School{}); err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBUsers() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("././database/Users.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(models.DataUser{}); err != nil {
		return nil, err
	}
	return db, nil
}
