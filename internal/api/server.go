package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router   *chi.Mux
	pgxClose io.Closer
	cfg      ServerConfig
}

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(router *chi.Mux, cfg ServerConfig) *Server {
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) Start() error {
	srv := &http.Server{
		Addr:         s.cfg.Addr,
		Handler:      s.router,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
		IdleTimeout:  s.cfg.IdleTimeout,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	errCh := make(chan error)

	go func() {
		log.Printf("🚀 starting on %s", srv.Addr)
		log.Printf("📁 docs available at http://%s%s/",
			srv.Addr, SwaggerPath)

		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("starting server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-stop:
		s.gracefulShutdown(srv)

		if s.pgxClose != nil {
			log.Printf("closing postgres connection")
			err := s.pgxClose.Close()
			if err != nil {
				log.Printf("error closing postgres connection: %v", err)
			}
		}
	}

	return nil
}

func (s *Server) gracefulShutdown(server *http.Server) {
	log.Println("🛑 shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("error shutting down server: %v", err)
		return
	}

	log.Println("✅ server stopped gracefully")
}

func (s *Server) SetPostgresCloser(closer io.Closer) {
	s.pgxClose = closer
}
