package http

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	Start()
	Stop() error
	Notify() <-chan error
}

type server struct {
	httpServer http.Server
	notify     chan error
}

func NewServer(addr string, router http.Handler) Server {
	return &server{
		httpServer: http.Server{Addr: addr, Handler: router},
		notify:     make(chan error),
	}
}

func (server *server) Start() {
	go func() {
		server.notify <- server.httpServer.ListenAndServe()
		close(server.notify)
	}()
}

func (server *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return server.httpServer.Shutdown(ctx)
}

func (server *server) Notify() <-chan error {
	return server.notify

}
