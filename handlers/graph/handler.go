package graph

import (
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
		utils.APIResponse(ctx, err.Error(), 402, http.MethodPost, nil)
		return
	}

	resultRegister, errRegister := h.service.GetGraph(&input)

	switch errRegister {
	case "VALIDATE_FAILED_402":
		utils.APIResponse(ctx, "Start point or goal point is in obstacles", 402, http.MethodPost, nil)
		return
	default:
		utils.APIResponse(ctx, "Get graph OK!", http.StatusOK, http.MethodPost, resultRegister)
		return
	}
}
