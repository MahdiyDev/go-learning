package lib

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	DATABASE_URL string
	db           *sqlx.DB
}

func (d *DatabaseConnection) Setup() error {
	db, err := sqlx.Connect("postgres", d.DATABASE_URL)
	d.db = db
	return err
}

func (d DatabaseConnection) FetchAll(
	dest interface{},
	query string,
	args ...interface{},
) error {
	err := d.db.Select(dest, query, args...)
	return err
}

func (d DatabaseConnection) FetchOne(
	dest interface{},
	query string,
	args ...interface{},
) error {
	err := d.db.Get(dest, query, args...)
	return err
}
