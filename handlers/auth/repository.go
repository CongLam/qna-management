package auth

import (
	model "qna-management/models"
	"qna-management/utils"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
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

func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var users model.EntityUsers
	db := r.db.Model(&users)
	var ErrorCode string

	checkUserAccount := db.Where("email = ?", input.Email).Find(&users)
	spew.Dump(input)
	if checkUserAccount.RowsAffected < 1 {
		ErrorCode = "USER_NOT_FOUND_404"
		return &users, ErrorCode
	}

	if !users.Active {
		ErrorCode = "LOGIN_NOT_ACTIVE_403"
		return &users, ErrorCode
	}

	comparePassword := utils.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		ErrorCode = "LOGIN_WRONG_PASSWORD_403"
		return &users, ErrorCode
	}
	return &users, ErrorCode
}
