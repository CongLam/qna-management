package login

import "qna-management/models"

type service struct {
	repository Repository
}

type Service interface {
	LoginService(input *InputLogin) (*models.EntityUsers, string)
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*models.EntityUsers, string) {
	user := models.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
