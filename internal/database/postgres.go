package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, url string) *pgxpool.Pool {
	cfg, err := pgxpool.ParseConfig(url)
    if err != nil {
        log.Fatal("[ERR] Unable to connect to database: ", err)
    }

    cfg.MaxConns = 20
    cfg.MinConns = 5
    cfg.MaxConnLifetime = time.Hour

    DB, err := pgxpool.NewWithConfig(ctx, cfg)

    if err != nil {
		log.Fatal("[ERR] Unable to connect to database: ", err)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatal("[ERR] Could not ping the database: ", err)
	}
	fmt.Println("[DEBUG] Connected to the database!")

    return DB
}
