package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func httpHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Test default"))
// }

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

func main() {
	var err error

	r := gin.Default()

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.LoadHTMLGlob("modern/templates/*.html")

	r.GET("/", index)
	r.GET("/:path", pathURL)

	// r.StaticFile("/index.html", "modern/templates/index.html")
	// r.StaticFile("/about.html", "modern/templates/about.html")
	// r.StaticFile("/contact.html", "modern/templates/contact.html")
	// r.StaticFile("/faq.html", "modern/templates/faq.html")
	// r.StaticFile("/pricing.html", "modern/templates/pricing.html")
	// r.StaticFile("/blog-home.html", "modern/templates/blog-home.html")
	// r.StaticFile("/blog-post.html", "modern/templates/blog-post.html")
	// r.StaticFile("/portfolio-item.html", "modern/templates/portfolio-item.html")
	// r.StaticFile("/portfolio-overview.html", "modern/templates/portfolio-overview.html")

	r.Static("/assets", "modern/assets")
	r.Static("/css", "modern/css")
	r.Static("/js", "modern/js")

	err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
