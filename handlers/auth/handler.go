package auth

import (
	"fmt"
	"net/http"

	"qna-management/utils"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func InitHandlerAuth(service Service) *handler {
	return &handler{service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input InputRegister
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		return
	}

	resultRegister, errRegister := h.service.RegisterService(&input)
	fmt.Println(resultRegister)

	switch errRegister {
	case "REGISTER_CONFLICT_409":
		utils.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		utils.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		// accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		// accessToken, errToken := utils.Sign(accessTokenData, utils.GodotEnv("JWT_SECRET"), 60)

		// if errToken != nil {
		// 	defer logrus.Error(errToken.Error())
		// 	utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		// 	return
		// }
		//
		// _, errSendMail := utils.SendGridMail(resultRegister.Fullname, resultRegister.Email, "Activation Account", "template_register", accessToken)
		//
		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	utils.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
		// 	return
		// }

		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
		return
	}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input InputLogin

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)
	switch errLogin {
	case "USER_NOT_FOUND_404":
		utils.APIResponse(ctx, "Login failed! User not found.", http.StatusForbidden, http.MethodPost, nil)
		return
	case "LOGIN_NOT_ACTIVE_403":
		utils.APIResponse(ctx, "Login failed! User not active", http.StatusForbidden, http.MethodPost, nil)
		return
	case "LOGIN_WRONG_PASSWORD_403":
		utils.APIResponse(ctx, "Login failed! Recheck your email or password", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, resultLogin)
		return
	}
}
