package routes

import (
	"bufio"
	"path/filepath"
	//"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	//"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MarkdownPayload struct {
	title string
	markdown_string string
}

type MetaDataPayload struct {
	description string
	date string
}

type ArticleCardPayload struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
}

var ARTICLE_DIR = "../articles"

func GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parent_dir = ARTICLE_DIR
		entries, err := os.ReadDir(parent_dir)
		var output []ArticleCardPayload

		if err != nil {
			log.Println("Error: ", err)
			c.String(http.StatusNotFound, "No Articles Found")
			c.Abort()
		}

		for _, entrie := range entries {
			log.Println(entrie.Name())
			description, err := getArticleDescription(entrie.Name())
			if err != nil {
				log.Println("Error: ", err)
			}
			imageUrl := fmt.Sprintf("/api/images?article=%s&image=%s", entrie.Name(), "thumbnail")
			newCard := ArticleCardPayload {
				Title: entrie.Name(),
				Description: description,
				ImageUrl: imageUrl,
			}

			output = append(output, newCard)
		}
		c.JSON(http.StatusOK, output);

	}
}

func GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		file, err := os.ReadFile(ARTICLE_DIR + "/" + name + "/" + name + ".md")

		if err != nil {
			log.Println("Error: ", err)
			c.String(http.StatusNotFound, "Article Body Not Found")
			c.Abort()
			return
		}

		c.Header("Content-Type", "text/markdown")
		c.String(http.StatusOK, string(file))

	}
}

func GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL.Query().Get("article"))
		article := c.Query("article")
		image := c.Query("image")

		if article == "" || image == "" {
			c.String(http.StatusBadRequest, "Missing Lookup Commands")
			c.Abort()
			return;
		}

		entries, err := os.ReadDir(ARTICLE_DIR + "/" + article);

		if err != nil {
			log.Println("Error: ", err)
			c.String(http.StatusNotFound, "Article Not Found")
			c.Abort()
			return
		}

		for _, entrie := range entries {
			if strings.Contains(entrie.Name(), image) {
				img, err := os.ReadFile(ARTICLE_DIR + "/" + article + "/" + entrie.Name())
				if err != nil {
					log.Println("Image Not Openable")
					c.String(http.StatusInternalServerError, "Image Not Availible")
					c.Abort()
					return;
				}
				contentType := ""
				ext := filepath.Ext(ARTICLE_DIR + "/" + article + "/" + entrie.Name())
				switch ext {
				case ".jpg", ".jpeg":
					contentType = "image/jpeg"
				case ".png":
					contentType = "image/png"
				case ".gif":
					contentType = "image/gif"
				default:
					contentType = "application/octet-stream" // Default if type can't be determined
				}
				log.Println("Extension: ", ext)
				c.Data(http.StatusOK, contentType, img)
				return
			}
		}
		log.Println("Image Not Found")
		c.String(http.StatusNotFound, "Image Not Found")
		c.Abort()


	}
}

func getArticleDescription(article string) (string, error) {
	file, err := os.Open(ARTICLE_DIR + "/" + article + "/meta.md")

	if err != nil {return "", err}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, '#') {
			continue
		}
		if strings.Contains(line, "Description: ") {
			log.Println(line)
			output := strings.Split(line, "Description: ")
			log.Println(output)
			return output[1], nil
		}
	}

	return "", errors.New("No Description Found in Metadata File")
}

