package handlers

import (
	"fmt"
	"go-web/config"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func indexHTMLPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})

	return
}

func dynamicHTMLPage(templateFiles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var path, pathHTML string
		path = c.Param("path")
		pathHTML = path + ".html"

		fileExists := containsFile(templateFiles, pathHTML)

		if !fileExists {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})

			return
		}

		c.HTML(http.StatusOK, pathHTML, gin.H{})

		return
	}
}

func fetchTemplatefiles(templateDir string) ([]string, error) {
	const sliceLenght int = 32
	listFiles := make([]string, sliceLenght)

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return listFiles, fmt.Errorf("function fetchTemplatefiles failed with error: %w", err)
	}

	for index, file := range files {
		listFiles[index] = file.Name()
	}

	return listFiles, nil
}

func containsFile(fetchedTemplateFiles []string, requestedPathHTML string) bool {
	for _, file := range fetchedTemplateFiles {
		if requestedPathHTML == file {
			return true
		}
	}

	return false
}

func Handler() *gin.Engine {
	conf := config.FetchVars()

	r := gin.Default()

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	templateFiles, err := fetchTemplatefiles(conf.TemplateDir)
	if err != nil {
		log.Fatalf("Function Handler failed with error: %v", err)
	}

	r.LoadHTMLGlob(conf.TemplateDir + "*.html")
	r.Static("/assets", "modern/assets")
	r.Static("/css", "modern/css")
	r.Static("/js", "modern/js")

	r.GET("/", indexHTMLPage)
	r.GET("/:path", dynamicHTMLPage(templateFiles))

	return r
}
