package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	listenAddr string
)

func shutdownServer(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":7123", "server listen address")
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	//让quit监听os.Interrupt信号量
	signal.Notify(quit, os.Interrupt)

	fs := http.FileServer(http.Dir("fileServer/assets/"))

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server at %s...\n", listenAddr)
	server := &http.Server{Addr: listenAddr, Handler: router}

	go shutdownServer(server, logger, quit, done)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}
