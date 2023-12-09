package api

import (
	"github.com/gin-gonic/gin"

	_ "easy_travel/api/docs"
	"easy_travel/api/handler"
	"easy_travel/config"
	"easy_travel/storage"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	// Country ...
	r.POST("/country", handler.CreateCountry)
	r.GET("/country/:id", handler.GetByIDCountry)
	r.GET("/country", handler.GetListCountry)
	r.PUT("/country/:id", handler.UpdateCountry)
	r.DELETE("/country/:id", handler.DeleteCountry)

	//City
	r.POST("/city", handler.CreateCity)
	r.GET("/city/:id", handler.GetByIDCity)
	r.GET("/city", handler.GetListCity)
	r.PUT("/city/:id", handler.UpdateCity)
	r.DELETE("/city/:id", handler.DeleteCity)

	//Aeroport
	r.POST("/aeroport", handler.CreateAeroport)
	r.GET("/aeroport/:id", handler.GetByIDAeroport)
	r.GET("/aeroport", handler.GetListAeroport)
	r.PUT("/aeroport/:id", handler.UpdateAeroport)
	r.DELETE("/aeroport/:id", handler.DeleteAeroport)

	//User
	r.POST("/user",handler.CheckPasswordMiddleware(), handler.CreateUser)
	r.GET("/user/:id", handler.GetByIDUser)
	r.GET("/user",handler.CheckApiKeyTime(), handler.Login)
	r.PUT("/user",handler.CheckPasswordMiddleware(), handler.UpdateExpire)
	
	//UploadFile
	r.POST("/upload", handler.UploadFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
