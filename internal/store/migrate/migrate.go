package migrate

import (
	"errors"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store/model"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return errors.New("migrate user model failed")
	}
	return nil
}
