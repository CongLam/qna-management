package auth

import (
	"fmt"

	model "qna-management/models"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	fmt.Println("In repository: Starting........")

	var user model.EntityUsers
	tx := r.db.Begin()
	// db := r.db.Model(&user)

	var errorCode string

	checkUserAccount := tx.Where("email = ?", input.Email).Find(&user)
	spew.Dump(checkUserAccount.RowsAffected)
	if checkUserAccount.RowsAffected > 0 {
		errorCode = "REGISTER_CONFLICT_409"
		tx.Rollback()
		return &user, errorCode
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := tx.Debug().Create(&user)

	if addNewUser.Error != nil {
		errorCode = "USER_INPUT_INVALID_422"
		tx.Rollback()
		return &user, errorCode
	}

	tx.Commit()
	return &user, errorCode
}
