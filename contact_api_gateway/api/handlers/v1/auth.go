package v1

import (
	"contact_api_gateway/api/models"
	"contact_api_gateway/genproto/user_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @ID create_user
// @Router /auth/sign-up [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserModel true "user"
// @Success 201 {object} models.ResponseModel{data=string} "User data"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
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

// SignIn godoc
// @ID sign-in-user
// @Router /auth/sign-in [POST]
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.SigninInput true "user"
// @Success 201 {object} models.ResponseModel{data=string} "User data"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) SignIn(c *gin.Context) {
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
