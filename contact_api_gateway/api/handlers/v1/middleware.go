package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"contact_api_gateway/api/models"
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

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		models.NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userID", userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get("userID")
	if !ok {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
