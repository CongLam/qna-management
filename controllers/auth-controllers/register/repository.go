package register

import (
	model "qna-management/models"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var user model.EntityUsers
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&user)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "USER_CONFLICT_409"
		return &user, <-errorCode
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := db.Debug().Create(&user)

	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "USER_INPUT_INVALID_422"
		return &user, <-errorCode
	}

	return &user, <-errorCode
}
