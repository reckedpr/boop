package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
	"github.com/reckedpr/boop/internal/dir"
	"github.com/reckedpr/boop/internal/server"
)

// hai

func main() {
	args := cli.ParseArgs()

	r := server.InitGin()

	var listenMsg string

	isPiped, data := cli.ReadStdin()
	if isPiped {
		listenMsg = "serving from stdin"

		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, string(data))
		})
	} else {
		listenMsg = fmt.Sprintf("serving %s", args.Path)

		r.GET("/*filepath", func(c *gin.Context) {
			dir.HandlePath(c, args.Path)
		})
	}

	srv := server.InitServer(r, &args, listenMsg)

	server.InterruptHandler(srv, &args)
}
