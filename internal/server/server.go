package server

import (
	"fmt"
	"html/template"
	"net/http"
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

	r.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		c.Header("Pragma", "no-cache") // 4 the dinosaurs out there
		c.Header("Expires", "0")
		c.Next()
	})

	r.Use(BoopLogger())

	return r
}

func InitServer(r *gin.Engine, args *cli.CliArgs, msg string) *http.Server {
	cli.BoopInfo("initialising http server")
	itf := fmt.Sprintf("127.0.0.1:%d", args.Port)
	if args.Host {
		itf = fmt.Sprintf(":%d", args.Port)
	}

	srv := &http.Server{
		Addr:    itf,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			cli.BoopErr(err.Error())
		}
	}()

	cli.BoopLog("%s on port %d (ctrl+c to stop)", msg, args.Port)
	PrintInterfaces(args.Port, args.Host)

	return srv
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
