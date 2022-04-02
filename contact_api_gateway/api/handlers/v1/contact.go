package v1

import (
	"context"
	"errors"

	"net/http"

	"contact_api_gateway/api/models"
	"contact_api_gateway/genproto/contact_service"
	"contact_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// Create Contact godoc
// @Security ApiKeyAuth
// @ID create-contact
// @Router /v1/contact [POST]
// @Summary create contact
// @Description Create Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param contact body models.CreateContactModel true "contact"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateContact(c *gin.Context) {

	userID, err := getUserID(c)
	if err != nil {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

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
			Name:   input.Name,
			Phone:  input.Phone,
			UserId: userID,
		},
	)
	if !handleError(h.log, c, err, "error while creating contact") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", id)

}

// Get All Contact godoc
// @Security ApiKeyAuth
// @ID get-all-contact
// @Router /v1/contact [GET]
// @Summary get all contact
// @Description Get All Contact
// @Tags contact
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseModel{data=models.GetAllContactModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllContact(c *gin.Context) {
	var contacts models.GetAllContactModel

	userID, err := getUserID(c)
	if err != nil {
		return
	}

	resp, err := h.services.ContactService().GetAll(
		context.Background(),
		&contact_service.UserId{
			UserId: userID,
		},
	)

	if !handleError(h.log, c, err, "error while getting all contacts") {
		return
	}

	err = ParseToStruct(&contacts, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", contacts)
}

// Get Contact godoc
// @Security ApiKeyAuth
// @ID get-contact
// @Router /v1/contact/{contact_id} [GET]
// @Summary get contact
// @Description Get Contact
// @Tags contact
// @Accept json
// @Produce json
// @Param contact_id path string true "contact_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetContact(c *gin.Context) {
	var contact models.Contact
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	contact_id := c.Param("contact_id")

	if !util.IsValidUUID(contact_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid contact id", errors.New("contact id is not valid"))
		return
	}

	resp, err := h.services.ContactService().Get(
		context.Background(),
		&contact_service.ContactUserId{
			Id:     contact_id,
			UserId: userID,
		},
	)

	if !handleError(h.log, c, err, "error while getting contact") {
		return
	}

	err = ParseToStruct(&contact, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", contact)
}

// Update Contact godoc
// @Security ApiKeyAuth
// @ID update_contact
// @Router /v1/contact/{contact_id} [PUT]
// @Summary Update Contact
// @Description Update Contact by ID
// @Tags contact
// @Accept json
// @Produce json
// @Param contact_id path string true "contact_id"
// @Param contact body models.CreateContactModel true "contact"
// @Success 200 {object} models.ResponseModel{data=models.ContactUpdate} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateContact(c *gin.Context) {
	var status models.ContactUpdate
	var contact contact_service.Contact

	userID, err := getUserID(c)
	if err != nil {
		models.NewErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	contact_id := c.Param("contact_id")

	if !util.IsValidUUID(contact_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid position id", errors.New("contact id is not valid"))
		return
	}

	if err := c.BindJSON(&contact); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binging json", err)
		return
	}

	resp, err := h.services.ContactService().Update(
		context.Background(),
		&contact_service.Contact{
			Id:     contact_id,
			Name:   contact.Name,
			Phone:  contact.Phone,
			UserId: userID,
		},
	)

	if !handleError(h.log, c, err, "error while getting contact") {
		return
	}

	err = ParseToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", status)

}

/*
// Delete Contact godoc
// @ID delete_profession
// @Security ApiKeyAuth
// @Router /v1/contact/{contact_id} [DELETE]
// @Summary Delete Contact
// @Description Delete Contact by given ID
// @Tags profession
// @Accept json
// @Produce json
// @Param Id path string true "contact_id"
// @Success 200 {object} models.ResponseModel{data=models.NameModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// func (h *handlerV1) DeleteContact(c *gin.Context) {
// 	userID, err := getUserID(c)
// 	if err != nil {
// 		return
// 	}

// 	var id int
// 	id, err = strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		models.NewErrorResponce(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	err = h.services.ContactService().Delete(userID, id)
// 	if err != nil {
// 		models.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "ok",
// 	})
// }
*/
