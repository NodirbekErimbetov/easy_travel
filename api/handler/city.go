package handler

import (
	"database/sql"
	"net/http"

	"easy_travel/models"
	"easy_travel/pkg/helpers"

	"github.com/gin-gonic/gin"
)

// CreateCity godoc
// @ID create_city
// @Router /city [POST]
// @Summary Create City
// @Description Create City
// @Tags City
// @Accept json
// @Produce json
// @Param object body models.CreateCity true "CreateCityRequestBody"
// @Success 200 {object} Response{data=models.City} "CityBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCity(c *gin.Context) {

	var createCity models.CreateCity
	err := c.ShouldBindJSON(&createCity)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}
	resp, err := h.strg.City().Create(c.Request.Context(), &createCity)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCity godoc
// @ID get_by_id_city
// @Router /city/{id} [GET]
// @Summary Get By Id City
// @Description Get By Id City
// @Tags City
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.City} "GetListCityResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDCity(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := h.strg.City().GetByID(c.Request.Context(), &models.CityPrimaryKey{Id: id})
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

// GetListCity godoc
// @ID get_list_city
// @Router /city [GET]
// @Summary Get List City
// @Description Get List City
// @Tags City
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListCityResponse} "GetListCityResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCity(c *gin.Context) {

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

	resp, err := h.strg.City().GetList(c.Request.Context(), &models.GetListCityRequest{
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

// UpdateCity godoc
// @ID city_update
// @Router /city [PUT]
// @Summary CityUpdate
// @Description CityUpdate
// @Tags City
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} Response{data=models.UpdateCity} "GetListCityResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCity(c *gin.Context) {

	var updateCity models.UpdateCity

	err := c.ShouldBindJSON(&updateCity)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateCity.Id = id

	rowsAffected, err := h.strg.City().Update(c.Request.Context(), &updateCity)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.strg.City().GetByID(c.Request.Context(), &models.CityPrimaryKey{Id: updateCity.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteCity godoc
// @ID city_delete
// @Router /city/{id} [DELETE]
// @Summary CityDelete
// @Description CityDelete
// @Tags City
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.CityPrimaryKey} "Success"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCity(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	err := h.strg.City().Delete(c.Request.Context(), &models.CityPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
