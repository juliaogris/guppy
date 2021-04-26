package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/juliaogris/guppy/pkg/echo"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var version = "v0.0.0-unset"

type config struct {
	Address string           `help:"hostname:port" default:"localhost:9090"`
	Version kong.VersionFlag `short:"V" help:"Print version information" group:"Other:"`
}

func main() {
	cfg := &config{}
	_ = kong.Parse(cfg, kong.Vars{"version": version})
	fmt.Println("starting echo guide server on", cfg.Address)
	if err := run(cfg.Address); err != nil {
		log.Fatal(err)
	}
}
func run(addr string) error {
	s := grpc.NewServer()
	echoServer := &echo.Server{}
	echo.RegisterEchoServer(s, echoServer)
	reflection.Register(s)

	h := &http.Server{
		Addr:    addr,
		Handler: rootHandler(s),
	}
	if err := h.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve gRPC service: %w", err)
	}
	return nil
}

// From: https://github.com/philips/grpc-gateway-example/issues/22#issuecomment-490733965
// Use x/net/http2/h2c so we can have http2 cleartext connections.
func rootHandler(grpcServer http.Handler) http.Handler {
	hf := func(w http.ResponseWriter, r *http.Request) {
		var label string
		switch {
		case r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc"):
			label = "grpc"
			grpcServer.ServeHTTP(w, r)
		default:
			label = "error"
			http.Error(w, r.URL.Path+": not Implemented", http.StatusNotImplemented)
		}
		fmt.Printf("%-5s: %-4s %s %s\n", label, r.Method, r.URL.Path, r.RemoteAddr)
	}
	return h2c.NewHandler(http.HandlerFunc(hf), &http2.Server{})
}
