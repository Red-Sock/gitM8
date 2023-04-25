package postgres

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/jackc/pgx/v5"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/utils/closer"
)

func New(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, createConnectionString(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error checking connection to redis")
	}

	closer.Add(func() error {
		return conn.Close(ctx)
	})

	return conn, nil
}

func createConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.GetString(config.DataSourcesPostgresDBUser),
		cfg.GetString(config.DataSourcesPostgresDBPwd),
		cfg.GetString(config.DataSourcesPostgresDBHost),
		cfg.GetString(config.DataSourcesPostgresDBPort),
		cfg.GetString(config.DataSourcesPostgresDBName),
	)
}
