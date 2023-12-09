package handler

import (
	"database/sql"
	"net/http"

	"easy_travel/models"
	"easy_travel/pkg/helpers"

	"github.com/gin-gonic/gin"
)

// CreateAeroport godoc
// @ID create_aeroport
// @Router /aeroport [POST]
// @Summary Create Aeroport
// @Description Create Aeroport
// @Tags Aeroport
// @Accept json
// @Produce json
// @Param object body models.CreateAeroport true "CreateAeroportRequestBody"
// @Success 200 {object} Response{data=models.Aeroport} "AeroportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAeroport(c *gin.Context) {

	var createAeroport models.CreateAeroport
	err := c.ShouldBindJSON(&createAeroport)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}

	resp, err := h.strg.Aeroport().Create(c.Request.Context(),&createAeroport)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdAeroport godoc
// @ID get_by_id_aeroport
// @Router /aeroport/{id} [GET]
// @Summary Get By Id Aeroport
// @Description Get By Id Aeroport
// @Tags Aeroport
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Aeroport} "GetListAeroportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDAeroport(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := h.strg.Aeroport().GetByID(c.Request.Context(), &models.AeroportPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListAeroport godoc
// @ID get_list_aeroport
// @Router /aeroport [GET]
// @Summary Get List Aeroport
// @Description Get List Aeroport
// @Tags Aeroport
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListAeroportResponse} "GetListAeroportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListAeroport(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query("limit"), 10)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query search")
		return
	}

	resp, err := h.strg.Aeroport().GetList(c.Request.Context(), &models.GetListAeroportRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateAeroport godoc
// @ID aeroport_update
// @Router /aeroport [PUT]
// @Summary AeroportUpdate
// @Description AeroportUpdate
// @Tags Aeroport
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} Response{data=models.UpdateAeroport} "GetListAeroportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateAeroport(c *gin.Context) {

	var updateAeroport models.UpdateAeroport

	err := c.ShouldBindJSON(&updateAeroport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateAeroport.Id = id

	rowsAffected, err := h.strg.Aeroport().Update(c.Request.Context(), &updateAeroport)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.strg.Aeroport().GetByID(c.Request.Context(), &models.AeroportPrimaryKey{Id: updateAeroport.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteAeroport godoc
// @ID aeroport_delete
// @Router /aeroport{id} [DELETE]
// @Summary AeroportDelete
// @Description AeroportDelete
// @Tags Aeroport
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.AeroportPrimaryKey} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server ErAeroport
func (h *Handler) DeleteAeroport(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := h.strg.Aeroport().Delete(c.Request.Context(), &models.AeroportPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
