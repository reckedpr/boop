package server

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
	"github.com/reckedpr/boop/internal/dir"
)

func InitGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	tpl := template.Must(template.New("dirlist").Parse(dir.HtmlTemplate))
	r.SetHTMLTemplate(tpl)

	return r
}

func InitServer(r *gin.Engine, args *cli.CliArgs, msg string) *http.Server {
	itf := fmt.Sprintf("127.0.0.1:%d", args.Port)
	if args.Host {
		itf = fmt.Sprintf(":%d", args.Port)
	}

	srv := &http.Server{
		Addr:    itf,
		Handler: r,
	}

	go func() {
		cli.BoopLog("%s on port %d (ctrl+c to stop)", msg, args.Port)
		PrintInterfaces(args.Port, args.Host)

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cli.BoopLog("error lol: %s\n", err)
			os.Exit(1)
		}
	}()

	return srv
}

func Shutdown(srv *http.Server, reason string) {
	cli.BoopLogNl("shutting down: %s", reason)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		cli.BoopLog("shutdown forcefully.. %s", err)
	}
}

func CatchInterrupt() <-chan os.Signal {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	return channel
}

func InterruptHandler(srv *http.Server, args *cli.CliArgs) {
	var expiry <-chan time.Time

	if args.Time > 0 {
		expiry = time.After(time.Duration(args.Time) * time.Minute)

		cli.BoopLog("expiring in %v minutes", args.Time)
	}

	interrupt := CatchInterrupt()

	select {
	case <-expiry:
		Shutdown(srv, "timer expired")
	case <-interrupt:
		Shutdown(srv, "caught interrupt")
	}
}
