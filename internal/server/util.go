package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/boop/internal/cli"
)

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

func BoopLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		latency := time.Since(start)
		ms := float64(latency) / float64(time.Millisecond)
		status := c.Writer.Status()

		s := fmt.Sprintf("%s %s (%.2fms)", c.Request.Method, c.Request.URL.Path, ms)
		cli.BoopHttp(status, s)
	}
}
