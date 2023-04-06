package postgres

import (
	"context"
	"fmt"

	"gitM8/internal/config"
)

func New(ctx context.Context, cfg *config.Config) (interface{}, error) { //(*pgx.Conn, error) {
	//conn, err := pgx.Connect(ctx, createConnectionString(cfg))
	//if err != nil {
	//	return nil, errors.Wrap(err, "error checking connection to redis")
	//}
	//
	//closer.Add(func() error {
	//	return conn.Close(ctx)
	//})

	//return conn, nil
	return nil, nil
}

func createConnectionString(cfg *config.Config) string {
	//"user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca"
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.GetString(config.DataSourcesPostgresUser),
		cfg.GetString(config.DataSourcesPostgresPwd),
		cfg.GetString(config.DataSourcesPostgresHost),
		cfg.GetString(config.DataSourcesPostgresPort),
		cfg.GetString(config.DataSourcesPostgresName),
	)
}
