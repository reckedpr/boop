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

// do not read the strings, just dont read the strin pls oki thankyu

func main() {
	pathFlag := pflag.StringP("path", "p", ".", "file or dirrr to uhhhhh ya")
	pflag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	tpl := template.Must(template.New("dirlist").Parse(web.HtmlTemplate))

	r.SetHTMLTemplate(tpl)

	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeCharDevice == 0 {
		data, _ := io.ReadAll(os.Stdin)
		//fmt.Println("raw the stdin cum in me, get me boypregnant ;3 mmhhhffgh")
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		//fmt.Println("no stdin input, using root or pathy wathy >_<")
		r.GET("/*filepath", func(c *gin.Context) {
			web.DisplayDir(c, *pathFlag)
		})
		fmt.Printf("boop serving %s\n", *pathFlag)
	}
	r.Run()
}
