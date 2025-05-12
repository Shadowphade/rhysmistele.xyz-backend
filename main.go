package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"rhysmistele.xyz/backend/routes"
)

func main() {

	router := gin.Default()

	router.Static("/static", "../site-build/static")
	router.NoRoute(func (c * gin.Context) {
		c.File("../site-build/index.html")
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
