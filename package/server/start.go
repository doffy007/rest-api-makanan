package server

import (
	"context"
	"fmt"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/doffy007/rest-api-makanan/config"
	"github.com/doffy007/rest-api-makanan/internal/middleware"
	"github.com/doffy007/rest-api-makanan/internal/router"
	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
}

type server struct {
	ctx   context.Context
	conf  *config.Config
	route *mux.Router
}

// Start implement.
func (s *server) Start() error {
	s.provider()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ConfigureApp.Port),
		Handler: s.route,
	}

	go func() {
		log.Printf("Server starting on port :%d\n", config.ConfigureApp.Port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return s.handleShutdown(server)
}

func (s *server) provider() {
	s.route.Use(middleware.LogMW)
	router.Register(s.ctx, s.conf, s.route).All()
}

func NewApp(ctx context.Context) Server {
	return &server{
		ctx:   ctx,
		route: mux.NewRouter(),
	}
}

func (s server) handleShutdown(srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}
