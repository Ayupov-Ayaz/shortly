//go:build tools

package tools

import (
	// migrate
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// oapi-codegen
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	// mockgen
	_ "go.uber.org/mock/mockgen"
)
