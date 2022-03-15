package v1

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Profession godoc
// @ID create-profession
// @Router /v1/profession [POST]
// @Summary create profession
// @Description Create Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body models.CreateProfessionModel true "profession"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfession(c *gin.Context) {
	var profession models.CreateProfessionModel

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.ProfessionService().Create(
		context.Background(),
		&position_service.Profession{
			Name: profession.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating profession") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

func (h *handlerV1) CreateContact(c *gin.Context) {

	userID, err := getUserID(c)

	var input contact.DefaultContact
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" || input.Phone == "" {
		newErrorResponce(c, http.StatusBadRequest, "invalid name or phone contact")
		return
	}

	id, err := h.service.Contacts.Create(input.Name, input.Phone, userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
