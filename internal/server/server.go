package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func CatchInterrupt() <-chan os.Signal {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	return channel
}

func PrintAddress(addr string, port int) {
	s := fmt.Sprintf("http://%s:%d", addr, port)
	s = cli.Colorise(s, cli.FgCyan)
	fmt.Printf("  > %s\n", s)
}

func PrintInterfaces(port int, host bool) {

	fmt.Printf("\n")
	PrintAddress("localhost", port)

	if !host {
		return
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.To4() == nil || ip.IsLoopback() {
				continue
			}

			PrintAddress(ip.String(), port)
		}
	}
}
