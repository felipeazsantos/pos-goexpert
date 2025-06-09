package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second) // Simulate a long-running request
		w.Write([]byte("Hello, World!"))
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start: ", err)
		}
	}()

	notifySignal := make(chan os.Signal, 1)
	signal.Notify(notifySignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-notifySignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server stopped")
}
