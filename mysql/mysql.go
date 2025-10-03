package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySql struct {
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime int

	Builder squirrel.StatementBuilderType
	DB      *sql.DB
}

func New(dsn string, opts ...Option) (*MySql, error) {
	m := &MySql{
		maxOpenConns:    defaultMaxOpenConns,
		maxIdleConns:    defaultMaxIdleConns,
		connMaxLifetime: defaultConnMaxLifetime,

		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
	}

	for _, opt := range opts {
		opt(m)
	}

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("mysqldb - New - sql.Open: %w", err)
	}

	dbConn.SetMaxIdleConns(m.maxIdleConns)
	dbConn.SetMaxOpenConns(m.maxOpenConns)
	dbConn.SetConnMaxLifetime(time.Duration(m.connMaxLifetime) * time.Minute)

	m.DB = dbConn
	return m, nil
}

func (m *MySql) Ping() error {
	if m.DB == nil {
		return errors.New("mysqldb - Ping - m.DB is nil")
	}

	err := m.DB.Ping()
	if err != nil {
		return fmt.Errorf("mysqldb - Ping - m.DB.Ping: %w", err)
	}

	return nil
}

func (m *MySql) Close() error {
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return fmt.Errorf("mysqldb - Close - m.DB.Close: %w", err)
		}
	}

	return nil
}
