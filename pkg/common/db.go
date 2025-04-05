package common

import (
	"github.com/alynlin/myapi/pkg/internal/repository/model"
	"gorm.io/gorm"
)

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	_ = sqlDB.Close()
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{})
}
