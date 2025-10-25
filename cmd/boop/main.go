package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/args"
	"github.com/reckedpr/boop/internal/web"
)

// hai

func initGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	tpl := template.Must(template.New("dirlist").Parse(web.HtmlTemplate))
	r.SetHTMLTemplate(tpl)

	return r
}

func main() {
	argPath := args.ParsePath()

	r := initGin()

	isPiped, data := args.ReadStdin()
	if isPiped {
		fmt.Println("boop serving from stdin")

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		fmt.Printf("boop serving %s\n", argPath)

		r.GET("/*filepath", func(c *gin.Context) {
			web.DisplayDir(c, argPath)
		})
	}

	r.Run()
}
