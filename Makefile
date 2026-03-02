run-development:
	go run ./cmd/api/main.go

run-production:
	go run ./cmd/api/main.go -env=production

generate-api:
	@echo "Generating API code from OpenAPI spec ..."
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		-config ./api/config.yaml \
		./api/openapi.yaml
	@echo "✅ API code generated"
