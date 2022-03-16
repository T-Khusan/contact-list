package v1

import (
	"context"

	"net/http"

	"contact_api_gateway/api/models"
	"contact_api_gateway/genproto/contact_service"

	"github.com/gin-gonic/gin"
)

// Create Contact godoc
// @ID create-contact
// @Router /v1/contact [POST]
// @Summary create contact
// @Description Create Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param profession body models.CreateContactModel true "contact"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateContact(c *gin.Context) {

	userID, err := getUserID(c)

	var input models.CreateContactModel

	if err := c.BindJSON(&input); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}
	if input.Name == "" || input.Phone == "" {
		models.NewErrorResponce(c, http.StatusBadRequest, "invalid name or phone contact")
		return
	}

	id, err := h.services.ContactService().Create(
		context.Background(),
		&contact_service.Contact{
			Name: input.Name,
			Phone: input.Phone,
			Id: userID,
		},
	)
	if !handleError(h.log, c, err, "error while creating contact") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", id)

}
