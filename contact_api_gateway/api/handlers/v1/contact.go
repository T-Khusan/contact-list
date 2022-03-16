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
			Name:  input.Name,
			Phone: input.Phone,
			Id:    userID,
		},
	)
	if !handleError(h.log, c, err, "error while creating contact") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", id)

}

/*
func (h *handlerV1) GetContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	mycontact, err := h.services.ContactService().GetByID(userID, id)
	output := contact.DefaultContact{
		Name:  mycontact.Name,
		Phone: mycontact.Phone,
	}
	c.JSON(http.StatusOK, output)
}

func (h *handlerV1) GetAllContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	contacts, err := h.services.ContactService().GetAll(userID)
	if err != nil {
		models.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	var output []allContact
	for _, value := range contacts {
		output = append(output, allContact{
			value.ID,
			value.Name,
			value.Phone,
		})
	}

	c.JSON(http.StatusOK, output)
}

func (h *handlerV1) UpdateContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input contact.DefaultContact
	if err := c.BindJSON(&input); err != nil {
		models.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ContactService().Update(userID, id, input)
	if err != nil {
		models.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *handlerV1) DeleteContact(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.ContactService().Delete(userID, id)
	if err != nil {
		models.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
*/
