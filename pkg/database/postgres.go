package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/igorscandido/go-transactions-api/pkg/configs"
	_ "github.com/lib/pq"
)

type postgresAdapter struct {
	conn *sql.DB
}

func NewPostgresAdapter(configs *configs.Configs) (*postgresAdapter, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		configs.Database.Driver,
		configs.Database.User,
		configs.Database.Password,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.DBName)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	return &postgresAdapter{conn: conn}, nil
}

func (p *postgresAdapter) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.conn.QueryContext(ctx, query, args...)
}

func (p *postgresAdapter) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.conn.ExecContext(ctx, query, args...)
}

func (p *postgresAdapter) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.conn.QueryRowContext(ctx, query, args...)
}

func (p *postgresAdapter) Close() error {
	return p.conn.Close()
}
