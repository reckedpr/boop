package server

import (
	"context"
	"net/http"
	"time"

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
