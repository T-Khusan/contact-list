package v1

import (
	"contact_api_gateway/api/models"
	"net/http"

	// "github.com/Mirobidjon/contact-list"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) signUp(c *gin.Context) {
	var user models.UserModel

	if err := c.BindJSON(&user); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	id, err := h.services.ContactService().CreateUser(user)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signinInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var user signinInput

	if err := c.BindJSON(&user); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
