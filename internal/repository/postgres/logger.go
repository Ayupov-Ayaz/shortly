package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
)

type Logger struct {
	// todo: zerolog.Logger
}

func (l *Logger) Log(
	ctx context.Context,
	level tracelog.LogLevel,
	msg string,
	data map[string]any,
) {
	// logging logic here, e.g. using zerolog
}
