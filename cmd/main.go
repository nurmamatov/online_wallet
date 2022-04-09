package main

import (
	"fmt"
	"log"
	"online_wallet/api"
	"online_wallet/config"
	"online_wallet/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	server := api.NewRouter(
		api.Option{
			Conf: cfg,
			Repo: storage.NewStoragePg(db),
		},
	)
	if err := server.Run(cfg.HttpPort); err != nil {
		log.Fatal("failed to run http server")
		panic(err)
	}

	defer db.Close()
}
