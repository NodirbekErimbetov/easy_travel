package handler

import (
	"database/sql"
	"net/http"

	"easy_travel/models"
	"easy_travel/pkg/helpers"

	"github.com/gin-gonic/gin"
)

// CreateCountry godoc
// @ID create_country
// @Router /country [POST]
// @Summary Create Country
// @Description Create Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CreateCountry true "CreateCountryRequestBody"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCountry(c *gin.Context) {

	var createCountry models.CreateCountry
	err := c.ShouldBindJSON(&createCountry)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}
	resp, err := h.strg.Country().Create(c.Request.Context(), &createCountry)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCountry godoc
// @ID get_by_id_country
// @Router /country/{id} [GET]
// @Summary Get By Id Country
// @Description Get By Id Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Country} "GetListCountrytResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDCountry(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := h.strg.Country().GetByID(c.Request.Context(), &models.CountryPrimaryKey{Id: id})
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

// GetListCountry godoc
// @ID get_list_country
// @Router /country [GET]
// @Summary Get List Country
// @Description Get List Country
// @Tags Country
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListCounrtyResponse}
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCountry(c *gin.Context) {

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

	resp, err := h.strg.Country().GetList(c.Request.Context(), &models.GetListCountryRequest{
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

// UpdateCountry godoc
// @ID country_update
// @Router /country [PUT]
// @Summary CountryUpdate
// @Description CountryUpdate
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.UpdateCountry true "UpdateCountryRequestBody"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCountry(c *gin.Context) {

	var updateCountry models.UpdateCountry
	err := c.ShouldBindJSON(&updateCountry)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateCountry.Id = id

	rowsAffected, err := h.strg.Country().UpdateCountry(c.Request.Context(), &updateCountry)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}
	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}
	resp, err := h.strg.Country().GetByID(c.Request.Context(), &models.CountryPrimaryKey{Id: updateCountry.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteCountry godoc
// @ID country_delete
// @Router /country/{id} [DELETE]
// @Summary CountryDelete
// @Description CountryDelete
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} Response{data=models.CountryPrimaryKey} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCountry(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := h.strg.Country().Delete(c.Request.Context(), &models.CountryPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}

