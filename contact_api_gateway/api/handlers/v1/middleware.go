package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"contact_api_gateway/api/models"
	"contact_api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) UserIdentify(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		models.NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// headerParts := strings.Split(header, " ")
	fmt.Println(header)
	// if headerParts[ == "" {
	// 	models.NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
	// 	return
	// }

	userID, err := h.services.UserService().ParseToken(
		context.Background(),
		&user_service.GetToken{
			Token: header,
		},
	)

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

	idStr, ok := id.(string)
	if !ok {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	return idStr, nil
}
