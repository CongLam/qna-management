package graph

import (
	"fmt"
	"net/http"

	"qna-management/utils"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func InitHandlerGraph(service Service) *handler {
	return &handler{service}
}

func (h *handler) GetGraph(ctx *gin.Context) {
	var input InputPoint
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		return
	}

	resultRegister, errRegister := h.service.GetGraph(&input)
	fmt.Println(resultRegister, errRegister)

	utils.APIResponse(ctx, "Get graph OK!", http.StatusOK, http.MethodPost, resultRegister)
	return
}
