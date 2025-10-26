package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
	"github.com/reckedpr/boop/internal/render"
	"github.com/reckedpr/boop/internal/server"
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
		listenMessage = "serving from stdin"

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		listenMessage = fmt.Sprintf("serving %s", args.Path)

		r.GET("/*filepath", func(c *gin.Context) {
			render.DisplayDir(c, args.Path)
		})
	}

	itf := fmt.Sprintf("127.0.0.1:%d", args.Port)
	if args.Host {
		itf = fmt.Sprintf(":%d", args.Port)
	}

	srv := &http.Server{
		Addr:    itf,
		Handler: r,
	}

	// start gorountine
	go func() {
		cli.BoopLog("%s on port %d (ctrl+c to stop)", listenMessage, args.Port)
		server.PrintInterfaces(args.Port, args.Host)

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cli.BoopLog("error lol: %s\n", err)
		}
	}()

	// timer stuff
	var expiry <-chan time.Time

	if args.Time > 0 {
		expiry = time.After(time.Duration(args.Time) * time.Minute)

		cli.BoopLog("stopping after %v minutes", args.Time)
	}

	interrupt := server.CatchInterrupt()

	select {
	case <-expiry:
		server.Shutdown(srv, "timer expired")
	case <-interrupt:
		server.Shutdown(srv, "caught interrupt")
	}
}
