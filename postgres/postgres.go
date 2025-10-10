package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type postgres struct {
	maxPoolSize  int32
	connAttempts int
	connTimeout  time.Duration

	pool       PgxPool
	sqlBuilder squirrel.StatementBuilderType
}

func New(uri string, opts ...Option) (DbEngine, error) {
	pg := &postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(pg)
	}

	poolConfig, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ParseConfig: %w", err)
	}
	poolConfig.MaxConns = pg.maxPoolSize

	for pg.connAttempts > 0 {
		pg.pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			err = pg.pool.Ping(context.Background())
			if err == nil {
				break
			}
		}

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if pg.connAttempts == 0 {
		return nil, fmt.Errorf("postgres - New - Connect postgres database: %w", err)
	}

	pg.sqlBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return pg, nil
}

func (pg *postgres) SqlBuilder() squirrel.StatementBuilderType {
	return pg.sqlBuilder
}

func (pg *postgres) PgxPool() PgxPool {
	return pg.pool
}

func (pg *postgres) Close() {
	if pg.pool != nil {
		pg.pool.Close()
	}
}
