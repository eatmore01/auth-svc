package postgresql

import (
	"auth/service/pkg/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageConfig struct {
	PgUser, PgDB, PgPass, PgPort string
}

type Client *pgxpool.Pool

func NewClient(ctx context.Context, sc StorageConfig) (pool *pgxpool.Pool, err error) {
	connectTimeOut := 5 * time.Second
	dsn := fmt.Sprintf("postgres://%s:%s@0.0.0.0:%s/%s", sc.PgUser, sc.PgPass, sc.PgPort, sc.PgDB)
	tries := 5

	utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), connectTimeOut)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			log.Fatalln("failed to connect to database: ", err.Error())
			return err
		}

		return nil
	}, tries, connectTimeOut)

	if err != nil {
		log.Fatal("failed to connect to database: ", err.Error())
		return nil, err
	}

	log.Print("connected to database successfully")
	return pool, nil
}
