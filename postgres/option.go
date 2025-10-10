package postgres

import "time"

const (
	defaultMaxPoolSize  = 10
	defaultConnAttempts = 5
	defaultConnTimeout  = time.Second
)

type Option func(*postgres)

func WithMaxPoolSize(maxPoolSize int32) Option {
	return func(p *postgres) {
		p.maxPoolSize = maxPoolSize
	}
}

func WithConnAttempts(attempts int) Option {
	return func(p *postgres) {
		p.connAttempts = attempts
	}
}

func WithConnTimeout(timeout time.Duration) Option {
	return func(p *postgres) {
		p.connTimeout = timeout
	}
}
