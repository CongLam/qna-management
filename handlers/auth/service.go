package auth

import (
	"fmt"

	"qna-management/models"
)

type service struct {
	repository Repository
}

type Service interface {
	RegisterService(input *InputRegister) (*models.EntityUsers, string)
	LoginService(input *InputLogin) (*models.EntityUsers, string)
}

func InitAuthService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*models.EntityUsers, string) {
	users := models.EntityUsers{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}
	fmt.Println("In service: Start call repo........")
	result, errRegister := s.repository.RegisterRepository(&users)
	fmt.Println("In service: End call repo........")

	return result, errRegister
}

func (s *service) LoginService(input *InputLogin) (*models.EntityUsers, string) {
	user := models.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	result, errLogin := s.repository.LoginRepository(&user)

	return result, errLogin
}
