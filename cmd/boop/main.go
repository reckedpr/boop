package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
	"github.com/reckedpr/boop/internal/render"
)

// hai

func initGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	tpl := template.Must(template.New("dirlist").Parse(render.HtmlTemplate))
	r.SetHTMLTemplate(tpl)

	return r
}

func main() {
	args := cli.ParseArgs()

	r := initGin()

	listenMessage := ""
	isPiped, data := cli.ReadStdin()
	if isPiped {
		listenMessage = "boop serving from stdin"

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		listenMessage = fmt.Sprintf("boop serving %s", args.Path)

		r.GET("/*filepath", func(c *gin.Context) {
			render.DisplayDir(c, args.Path)
		})
	}

	fmt.Printf("%s on port %d (ctrl+c to stop)\n", listenMessage, args.Port)

	itf := fmt.Sprintf(":%d", args.Port)
	r.Run(itf)
}
