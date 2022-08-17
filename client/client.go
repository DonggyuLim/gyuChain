package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GET  = "GET"
	POST = "POST"
)

type url string
type Documentaion struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(c *gin.Context) {
	document := []Documentaion{
		{
			URL:         url("/"),
			Method:      GET,
			Description: "Hello Documentation",
		},
	}
	c.JSON(http.StatusOK, document)
}

func Start(port string) {
	port = fmt.Sprintf(":%s", port)

	router := gin.Default()
	router.GET("/", documentation)

	log.Fatal(router.Run(port))

}
