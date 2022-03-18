package v1

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"contact_api_gateway/api/models"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) userIdentify(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		models.NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		models.NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.UserService().ParseToken(c context.Context, headerParts[1]) //h.service.Authorization
	if err != nil {
		models.NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userID", userID)
}

func getUserID(c *gin.Context) (string, error) {
	id, ok := c.Get("userID")
	if !ok {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	idInt, ok := id.(string)
	if !ok {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	return idInt, nil
}
