package handler

import (
	"easy_travel/config"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckPasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		password := c.GetHeader("API-KEY")
		if password != config.SecureApiKey {
			c.AbortWithError(http.StatusForbidden, errors.New("The request requires an user authentication."))
			return
		}

		c.Next()
	}
}

func (h *Handler) CheckApiKeyTime () gin.HandlerFunc {
	return func(c *gin.Context) {

	expireAt := config.ApiKeyExpiredAt
	timeNow := time.Now().Add(expireAt)

	if time.Now().After(timeNow){
		handleResponse(c,http.StatusForbidden,"Expired API key")
		return 
	}
	c.Next()
}
}