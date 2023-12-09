package handler

import (
	"easy_travel/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Upload File godoc
// @ID upload_file
// @Router /upload [POST]
// @Summary Upload file
// @Description Upload file
// @Tags File
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Param table_slug query string true "table_slug"
// @Success 200 {object} Response{} "File"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	data, err := file.Open()
	if err != nil {
		log.Println(err)
	}
	jsonFile, err := io.ReadAll(data)
	if err != nil {
		log.Println("Err while reading file: ", err)
	}
	fmt.Println(jsonFile)
	table := c.Query("table_slug")

	if table == "country" {
		var countries []models.CreateCountry
		err = json.Unmarshal(jsonFile, &countries)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, country := range countries {
			resp, err := h.strg.Country().Create(c.Request.Context(), &country)
			if err != nil {
				handleResponse(c, http.StatusInternalServerError, err)
			}
			handleResponse(c, http.StatusOK, resp)
		}
	}
	if table == "city" {
		var cities []models.CreateCity
		err := json.Unmarshal(jsonFile, &cities)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, city := range cities {
			resp, err := h.strg.City().Create(c.Request.Context(), &city)
			if err != nil {
				handleResponse(c, http.StatusInternalServerError, err)
			}

			handleResponse(c, http.StatusOK, resp)
		}
	}

	if table == "aeroport" {
		var aeroports []models.CreateAeroport
		err := json.Unmarshal(jsonFile, &aeroports)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, aeroport := range aeroports {
			resp, err := h.strg.Aeroport().Create(c.Request.Context(), &aeroport)
			if err != nil {
				handleResponse(c, http.StatusInternalServerError, err)
			}

			handleResponse(c, http.StatusOK, resp)
		}
	}

}

// Upload file godo
// @Router /upload [post]
// @ID file_upload
// @Summary Upload file
// @Description Upload file
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param table_slug query string true "table_slug"
// @Param file formData file true "this is a test file"
// @Success 200 {object} Response{data=string} "ok"
// @Failure 400 {object} Response{data=string} "We need ID!!"
// @Failure 404 {object} Response{data=string} "Can not find file"
func (h *Handler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handleResponse(c, 404, "file not")
		return
	}

	dst, _ := os.Getwd()
	dst += "/" + uuid.New().String() + "_" + file.Filename
	c.SaveUploadedFile(file, dst)
	defer os.Remove(dst)

	f, err := os.Open(dst)
	fmt.Println(f)	
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var data []map[string]interface{}

	body, err := io.ReadAll(f)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.Upload().Upload(c.Request.Context(), &models.UploadRequest{
		TableSlug: c.Query("table_slug"),
		Data:      data,
	})

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, http.StatusOK, "successfull")
}
