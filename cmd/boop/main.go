package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	itf := fmt.Sprintf(":%d", args.Port)

	srv := &http.Server{
		Addr:    itf,
		Handler: r,
	}

	// start gorountine
	go func() {
		cli.BoopLog("%s on port %d (ctrl+c to stop)", listenMessage, args.Port)

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cli.BoopLog("error lol: %s\n", err)
		}
	}()

	// if time arg not default use tima, else normal
	// TODO make uh CatchInterrupt func to handle interrupts on both timer and non timer
	if args.Time > 0 {
		expiry := time.Duration(args.Time) * time.Minute
		cli.BoopLog("stopping after %v", expiry)

		<-time.After(expiry)

		server.Shutdown(srv, "timer expired !")

		os.Exit(0)
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		signal := <-c
		reason := fmt.Sprintf("caught signal %v", signal)

		server.Shutdown(srv, reason)
	}
}
