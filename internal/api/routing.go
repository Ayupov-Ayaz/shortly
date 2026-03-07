package api

import (
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	swagger "github.com/swaggo/http-swagger"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest/handler"
)

const (
	V1Prefix     = "/api/v1"
	LivenessPath = "/live"
	MetricsPath  = "/metrics"

	SwaggerPath = V1Prefix + "/swagger"
	ShortenPath = V1Prefix + "/shorten"
)

// configureLiveness - configure liveness endpoint for health checks
func configureLiveness(
	respWriter *handler.ResponseWriter,
	router chi.Router,
) {
	router.Get(LivenessPath,
		func(writer http.ResponseWriter, _ *http.Request) {
			respWriter.SendOk(writer, "alive")
		})
}

// configureSwagger - configure swagger endpoint for API documentation
func configureSwagger(
	respWriter *handler.ResponseWriter,
	router chi.Router,
) {
	router.Get(path.Join(SwaggerPath, "*"), swagger.Handler(
		swagger.URL(path.Join(SwaggerPath, "doc.json")),
	))

	router.Get(path.Join(SwaggerPath, "doc.json"),
		handler.NewSwagger(respWriter).GetSwaggerUI)
}

// configureShortener - configure shortener endpoints for URL shortening service
func configureShortener(
	respWriter *handler.ResponseWriter,
	router chi.Router,
	shortener shortener.Shortener,
) {
	gen.HandlerFromMuxWithBaseURL(
		handler.NewURLShortener(shortener, respWriter),
		router,
		V1Prefix,
	)

}
