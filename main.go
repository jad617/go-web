package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func index(c *gin.Context) {
	// c.File("modern/index.html")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func pathURL(c *gin.Context) {
	// c.File("modern/about.html")
	var path, pathURI string
	path = c.Param("path")
	pathURI = path + ".html"

	c.HTML(http.StatusOK, pathURI, gin.H{})
}

func handler() *gin.Engine {
	r := gin.Default()

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.LoadHTMLGlob("modern/templates/*.html")
	r.Static("/assets", "modern/assets")
	r.Static("/css", "modern/css")
	r.Static("/js", "modern/js")

	r.GET("/", index)
	r.GET("/:path", pathURL)

	return r
}

func main() {
	r := handler()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
