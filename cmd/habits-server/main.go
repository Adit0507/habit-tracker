package main

import (
	"habits/internal/server"
	"habits/internal/log"
	"os"
)

const port = 28710

func main() {
	lgr := log.New(os.Stdout)
	srv := server.New(lgr) 
	
	err := srv.ListenAndServe(port)
	if err != nil {
	lgr.Logf("Error while running the server: %s", err.Error())
	os.Exit(1) 
	}
}
