package v1

import (
	"contact_api_gateway/api/models"
	"contact_api_gateway/genproto/user_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) SignUp(c *gin.Context) {
	var user models.UserModel

	if err := c.BindJSON(&user); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	id, err := h.services.UserService().CreateUser(
		context.Background(),
		&user_service.User{
			Id:       user.ID,
			Name:     user.Name,
			Lastname: user.Lastname,
			Password: user.Password,
		},
	)
	if !handleError(h.log, c, err, "error while creating user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", id)
}

func (h *handlerV1) signIn(c *gin.Context) {
	var user models.SigninInput

	if err := c.BindJSON(&user); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	token, err := h.services.UserService().GenerateToken(
		context.Background(),
		&user_service.GetAllUserRequest{
			Password: user.Password,
			Name:     user.Name,
		},
	)
	if !handleError(h.log, c, err, "error while signing user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", token)
}
