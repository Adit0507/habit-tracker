package main

import (
	"context"
	"habits/internal/log"
	"habits/internal/repository"
	"habits/internal/server"
	"os/signal"
	"syscall"

	// "habits/internal/repository"
	"os"
)

const port = 28710

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	lgr := log.New(os.Stdout)
	db := repository.New(lgr)
	srv := server.New(db, lgr)

	err := srv.ListenAndServe(ctx, port)
	if err != nil {
		lgr.Logf("Error while running the server: %s", err.Error())
		os.Exit(1)
	}
}
