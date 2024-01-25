package loginAuth

import (
	"qna-management/models"

	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *models.EntityUsers) (*models.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}
