package db

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
)

func ConnectDb() (db *bun.DB, ctx context.Context) {
	dsn := "postgresql://mohamedmansour:8v7TLmRapRIxLlwl8Tct6Q@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/runny-rhino-2736.tribal?sslmode=verify-full"

	ctx = context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	return db, ctx
}
