package main

import (
	"time"

	server ".."
)

func main() {
	srv := server.Server{
		Addr:         ":8080",
		IdleTimeout:  10 * time.Second,
		MaxReadBytes: 1000,
	}
	go srv.ListenAndServe()
	time.Sleep(10 * time.Second)

	select {}
}
