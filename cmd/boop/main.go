package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/web"
	"github.com/spf13/pflag"
)

// hai

func main() {
	pathFlag := pflag.StringP("path", "p", "", "file or dirrr to uhhhhh ya")
	pflag.Parse()

	var servePath string
	if *pathFlag != "" {
		servePath = *pathFlag
	} else if len(pflag.Args()) > 0 {
		servePath = pflag.Args()[0]
	} else {
		servePath = "."
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	tpl := template.Must(template.New("dirlist").Parse(web.HtmlTemplate))

	r.SetHTMLTemplate(tpl)

	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeCharDevice == 0 {
		fmt.Println("boop serving from stdin")

		data, _ := io.ReadAll(os.Stdin)

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		fmt.Printf("boop serving %s\n", servePath)

		r.GET("/*filepath", func(c *gin.Context) {
			web.DisplayDir(c, servePath)
		})
	}
	r.Run()
}
