package main

import (
	"habits/internal/log"
	"habits/internal/repository"
	"habits/internal/server"

	// "habits/internal/repository"
	"os"
)

const port = 28710

func main() {
	lgr := log.New(os.Stdout)
	db := repository.New(lgr)
	srv := server.New(db, lgr) 
	
	err := srv.ListenAndServe(port)
	if err != nil {
	lgr.Logf("Error while running the server: %s", err.Error())
	os.Exit(1) 
	}
}
