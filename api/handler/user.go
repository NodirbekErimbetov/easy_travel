package handler

import (
	"context"
	"database/sql"
	"net/http"
	"easy_travel/config"
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
func (h *Handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.ApiKeyExpiredAt)
	defer cancel()

	resp, err := h.strg.User().Create(ctx, &createUser)
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
func (h *Handler) GetByIDUser(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	// ctx,cancel := context.WithTimeout(context.Background(), config.ApiKeyExpiredAt)
	// defer cancel()

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
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

func (h *Handler) Login(c *gin.Context) {

	var userLogin models.User
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(400, "Shouldnim json err"+err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.ApiKeyExpiredAt)
	defer cancel()

	resp, err := h.strg.User().LoginUser(ctx, &models.LoginUser{Login: userLogin.Login, Password: userLogin.Password})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}

	if userLogin.Password != resp.Password {
		handleResponse(c, http.StatusBadRequest, "Login or Password wrong")
		return
	}
	if userLogin.Login != resp.Login {
		handleResponse(c, http.StatusBadRequest, "Login or Password wrong")
		return
	}

	handleResponse(c, http.StatusOK, resp)

}



func (h *Handler) UpdateExpire(c *gin.Context) {

	var userLogin models.UserPrimaryKey
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		handleResponse(c,http.StatusBadRequest,err.Error())
	}
	rows,err := h.strg.User().UpdateExpireTime(context.Background(),&userLogin)
	
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}
	if rows == 0 {
        handleResponse(c, http.StatusBadRequest, "no rows in result set")
        return
    }
	resp,err := h.strg.User().GetByID(context.Background(),&models.UserPrimaryKey{Id: userLogin.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError,err)
	}
	handleResponse(c,http.StatusAccepted,resp)
}
