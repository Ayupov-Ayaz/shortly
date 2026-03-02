package api

import (
	"fmt"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	swagger "github.com/swaggo/http-swagger"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/repository"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest/handler"
)

func Configure(
	router chi.Router,
	cfg *config.Config,
) error {
	respWriter := handler.NewResponseWriter()

	generator, err := id.NewSnowflakeGenerator(cfg.APP.Shortener.NodeID)
	if err != nil {
		return fmt.Errorf("creating snowflake generator: %w", err)
	}

	baseURL, err := cfg.Server.BaseURL(cfg.Env)
	if err != nil {
		return fmt.Errorf("parsing base url: %w", err)
	}

	shortenerSrv := shortener.New(repository.NewStubRepository(),
		generator,
		baseURL,
		cfg.APP.ShortURLsTTL(),
	)

	configureShorten(respWriter, router, shortenerSrv)
	configureLiveness(respWriter, router)
	configureSwagger(respWriter, router)

	return nil
}

func configureLiveness(
	respWriter *handler.ResponseWriter,
	router chi.Router,
) {
	router.Get(rest.LivenessPath,
		func(writer http.ResponseWriter, _ *http.Request) {
			respWriter.SendOk(writer, "alive")
		})
}

func configureSwagger(
	respWriter *handler.ResponseWriter,
	router chi.Router,
) {
	router.Get(path.Join(rest.SwaggerPath, "*"), swagger.Handler(
		swagger.URL(path.Join(rest.SwaggerPath, "doc.json")),
	))

	router.Get(path.Join(rest.SwaggerPath, "doc.json"),
		handler.NewSwagger(respWriter).GetSwaggerUI)
}

func configureShorten(
	respWriter *handler.ResponseWriter,
	router chi.Router,
	shortener shortener.Shortener,
) {
	gen.HandlerFromMuxWithBaseURL(
		handler.NewURLShortener(shortener, respWriter),
		router,
		rest.V1Prefix,
	)

}
