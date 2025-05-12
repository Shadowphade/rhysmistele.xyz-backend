package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"rhysmistele.xyz/backend/routes"
)

func main() {
	router := gin.Default()

	buildPath := filepath.Join("..", "rhysmistele.xyz-frontend", "dist")

	router.Static("/assets", filepath.Join(buildPath, "assets"))
	router.NoRoute(func (c * gin.Context) {
		c.File(filepath.Join(buildPath, "index.html"))
	})

	api := router.Group("/api")
	{
		api.GET("/articles/", routes.GetArticles())
		api.GET("/article/:id", routes.GetArticle())
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Running Server On http://localhost:%s\n", port)

	router.Run(":" + port)
}
