package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/HiogoPariz/files-notez/internal/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	//TODO SEPARAR MIDDLEWARE / ARRUMAR ORIGINS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("URL")}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Get user value
	router.GET("/:path", func(c *gin.Context) {
		path := c.Params.ByName("path")
		var file map[string]interface{}

		content, err := storage.ReadFile(path)
		if err != nil {
			c.AbortWithError(505, err)
			return
		}

		if err := json.Unmarshal([]byte(content), &file); err != nil {
			c.AbortWithError(505, err)
			return
		}

		c.JSON(http.StatusOK, &file)
	})

	router.POST("/:path", func(c *gin.Context) {
		path := c.Params.ByName("path")

		content, err := c.GetRawData()
		if err != nil {
			c.AbortWithError(505, err)
			return
		}

		if err := storage.WriteFile(path, string(content)); err != nil {
			c.AbortWithError(505, err)
			return
		}

		result := Message{Type: "success", Message: "created " + path, Code: 0}

		c.JSON(http.StatusOK, result)
	})

	router.DELETE("/:path", func(c *gin.Context) {
		path := c.Params.ByName("path")

		err := storage.DeleteFile(path)
		if err != nil {
			c.AbortWithError(505, err)
			return
		}

		result := Message{Type: "success", Message: "deleted " + path, Code: 0}

		c.JSON(http.StatusOK, result)
	})

	return router
}
