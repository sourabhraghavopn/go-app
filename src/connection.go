package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

func (app *App) loadConnection() Conn {
	dbName, exist := os.LookupEnv("DB_NAME")
	dbPort, exist := os.LookupEnv("DB_PORT")
	dbPassword, exist := os.LookupEnv("DB_PASSWORD")
	dbUserName, exist := os.LookupEnv("DB_USERNAME")

	if !exist {
		app.logger.Print("properties are missing")
		panic("properties are missing")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUserName, dbPassword, dbPort, dbName)
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

	return Conn{
		ctx: context.Background(),
		db:  db,
	}

}
