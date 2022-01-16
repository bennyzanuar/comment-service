package database

import (
	"github.com/bennyzanuar/comment-service/internal/comment"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&comment.Comment{})
	if err != nil {
		return err
	}
	return nil
}
