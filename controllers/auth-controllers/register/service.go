package register

import "qna-management/models"

type service struct {
	repository Repository
}

type Service interface {
	RegisterService(input *InputRegister) (*models.EntityUsers, string)
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*models.EntityUsers, string) {
	users := models.EntityUsers{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}

	result, errRegister := s.repository.RegisterRepository(&users)

	return result, errRegister
}
