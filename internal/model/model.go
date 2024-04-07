package model

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Pg *sqlx.DB // postgres
}

var db Database

func SetPostgresDB(pg *sqlx.DB) {
	db = Database{Pg: pg}
}

func (d *Database) NamedSelectContext(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	q, args, err := sqlx.BindNamed(sqlx.BindType(d.Pg.DriverName()), query, arg)
	if err != nil {
		return err
	}

	return d.Pg.SelectContext(ctx, dest, q, args...)
}

func (d *Database) NamedGetContext(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	q, args, err := sqlx.BindNamed(sqlx.BindType(d.Pg.DriverName()), query, arg)
	if err != nil {
		return err
	}

	return d.Pg.GetContext(ctx, dest, q, args...)
}

func (d *Database) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	q, args, err := sqlx.BindNamed(sqlx.BindType(d.Pg.DriverName()), query, arg)
	if err != nil {
		return nil, err
	}

	return d.Pg.ExecContext(ctx, q, args...)
}

func (d *Database) NamedExecContextReturnID(ctx context.Context, query string, arg interface{}, ID interface{}) error {
	q, args, err := sqlx.BindNamed(sqlx.BindType(d.Pg.DriverName()), query, arg)
	if err != nil {
		return err
	}

	return d.Pg.QueryRowx(q, args...).Scan(ID)
}

func (d *Database) NamedExecContextReturnObj(ctx context.Context, query string, arg interface{}, obj interface{}) error {
	q, args, err := sqlx.BindNamed(sqlx.BindType(d.Pg.DriverName()), query, arg)
	if err != nil {
		return err
	}

	return d.Pg.QueryRowx(q, args...).StructScan(obj)
}
