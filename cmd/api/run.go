package api

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/ayupov-ayaz/shortly/internal/api"
	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest"
)

const (
	prefix = "SHORTLY_"
)

func Run() error {
	env := parseEnvFlag()
	envFileName := choseEnvFile(env)

	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		return fmt.Errorf("parsing env: %w", err)
	}

	mux := rest.NewServer(rest.Config{
		Env:          env,
		LivenessPath: rest.LivenessPath,
		Timeout:      5 * time.Second,
	})

	err = api.Configure(mux, cfg)
	if err != nil {
		return fmt.Errorf("configuring api: %w", err)
	}

	err = start(makeServer(mux, cfg.Server))
	if err != nil {
		return err
	}

	return nil
}

func makeServer(router *chi.Mux, cfg config.Server) *http.Server {
	return &http.Server{
		Addr:         cfg.ListenAddr(),
		Handler:      router,
		ReadTimeout:  cfg.Timeout.Read,
		WriteTimeout: cfg.Timeout.Write,
		IdleTimeout:  cfg.Timeout.Idle,
	}
}

func start(server *http.Server) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	errCh := make(chan error)

	go func() {
		log.Printf("🚀 starting on %s", server.Addr)
		log.Printf("📁 docs available at http://%s%s/",
			server.Addr, rest.SwaggerPath)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("starting server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-stop:
		shutdown(server)
	}

	return nil
}

func shutdown(server *http.Server) {
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

func parseEnvFlag() string {
	env := flag.String("env", "development", "environment")
	flag.Parse()

	return *env
}

func choseEnvFile(env string) string {
	const (
		devEnvFile  = ".env.development"
		prodEnvFile = ".env.production"
	)

	switch env {
	case config.EnvDevelopment:
		return devEnvFile
	default:
		return prodEnvFile
	}
}
