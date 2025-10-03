package mysqldb

const (
	defaultMaxOpenConns    = 10
	defaultMaxIdleConns    = 10
	defaultConnMaxLifetime = 5
)

type Option func(*MySql)

func SetMaxOpenConns(n int) Option {
	return func(db *MySql) {
		db.maxOpenConns = n
	}
}

func SetMaxIdleConns(n int) Option {
	return func(db *MySql) {
		db.maxIdleConns = n
	}
}

func SetConnMaxLifetime(d int) Option {
	return func(db *MySql) {
		db.connMaxLifetime = d
	}
}
