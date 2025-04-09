package initializers

import (
	"pathipat/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User_test{})
}
