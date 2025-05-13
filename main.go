package main

import (
	"fmt"
	"os"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"rhysmistele.xyz/backend/routes"
)

func main() {
	router := gin.Default()
	router.Use((func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}

			c.Next()
	}))

	buildPath := filepath.Join("..", "rhysmistele.xyz-frontend", "dist")

	router.Static("/assets", filepath.Join(buildPath, "assets"))
	router.NoRoute(func (c * gin.Context) {
		c.File(filepath.Join(buildPath, "index.html"))
	})


	api := router.Group("/api")
	{
		api.GET("/articles", routes.GetArticles())
		api.GET("/article/:name", routes.GetArticle())
		api.GET("/images", routes.GetImage())
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Running Server On http://localhost:%s\n", port)

	router.Run(":" + port)
}
