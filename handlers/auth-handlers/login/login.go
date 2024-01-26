package login

import (
	loginAuth "qna-management/controllers/auth-controllers/login"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) LoginHandler(ctx *gin.Context) {

}
