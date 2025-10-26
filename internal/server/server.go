package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Shutdown(srv *http.Server, reason string) {
	fmt.Printf("\nboop shutting down: %s\n", reason)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("shutdown forcefully.. ", err)
	}
}
