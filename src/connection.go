package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func (app *App) loadDataSource() Conn {
	dbName, exist := os.LookupEnv("DB_NAME")
	dbPort, exist := os.LookupEnv("DB_PORT")
	dbPassword, exist := os.LookupEnv("DB_PASSWORD")
	dbUserName, exist := os.LookupEnv("DB_USERNAME")
	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		app.logger.Print("properties are missing")
		panic("properties are missing")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUserName, dbPassword,host, dbPort, dbName)
	fmt.Print(fmt.Sprintf("DSN : -|%s|- ",dsn))
	db := bun.NewDB(
		sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithDSN(dsn)),
		),
		pgdialect.New())
	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	db.RegisterModel((*UrlDetail)(nil))
	return Conn{
		ctx: context.Background(),
		db:  db,
	}

}

func (app *App) loadInMemorySqlLite() Conn {

	sqlite, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	sqlite.SetMaxOpenConns(1)

	db := bun.NewDB(sqlite, sqlitedialect.New())
	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
	db.RegisterModel((*UrlDetail)(nil))
	return Conn{
		ctx: context.Background(),
		db:  db,
	}

}
