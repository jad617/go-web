package main

import (
	"go-web/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var err error

const ctxKeyTemplateDir string = "templateDir"

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func pathURL(c *gin.Context) {
	var path, pathURI string
	path = c.Param("path")
	pathURI = path + ".html"

	templateDir := c.GetString(ctxKeyTemplateDir)

	templateFiles := fetchTemplatefiles(templateDir)

	validate := containsFile(templateFiles, pathURI)

	if validate {
		c.HTML(http.StatusOK, pathURI, gin.H{})
	} else {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	}
}

func fetchTemplatefiles(templateDir string) []string {
	const sliceLenght int = 32
	listFiles := make([]string, sliceLenght)

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatal(err)
	}

	for index, file := range files {
		listFiles[index] = file.Name()
	}

	return listFiles
}

func containsFile(list []string, str string) bool {
	for _, file := range list {
		if str == file {
			return true
		}
	}

	return false
}

func injectContext(ctxKey string, ctxValue string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Set("templateCTX", "modern/templates/")
		c.Set(ctxKey, ctxValue)
	}
}

func handler() *gin.Engine {
	conf := config.FetchVars()

	r := gin.Default()

	r.Use(injectContext(ctxKeyTemplateDir, conf.TemplateDir))

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.LoadHTMLGlob(conf.TemplateDir + "*.html")
	r.Static("/assets", "modern/assets")
	r.Static("/css", "modern/css")
	r.Static("/js", "modern/js")

	r.GET("/", index)
	r.GET("/:path", pathURL)

	return r
}

func init() {
	// For local development we load the .env file
	if os.Getenv("GO_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	r := handler()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
