package routes

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type MarkdownPayload struct {
	id int
	markdown_string string
}

func GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		entries, err := os.ReadDir("../articles")

		if err != nil {
			log.Println("Error: ", err)
			c.String(http.StatusNotFound, "No Articles Found")
			c.Abort()
		}

		for _, entrie := range entries {
			log.Println(entrie.Name())
		}

	}
}

func GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			log.Println("Error: ", err)
			c.String(http.StatusNotFound, "Article Not Found")
			c.Abort()
		}

		log.Println(id)
	}
}


func calculateCheckSum(inputFileName string) int {
	var output int;
	for character := range inputFileName {
		output = output + character
	}
	return output
}
