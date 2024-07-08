package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-storage/config"
	"product-storage/migrations"
	"product-storage/service"
	"product-storage/storage/postgres"
	"product-storage/transport"
	"product-storage/transport/middleware"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal("parse config failed: ", err)
	}
	log.Println("Config parsed")

	if err = migrations.MigrateUp(cfg.Db); err != nil {
		log.Fatal("migration failed: ", err)
	}
	log.Println("Migrations applied")

	s, err := postgres.NewStorage(cfg.Db)
	if err != nil {
		log.Fatal("storage initialization failed: ", err)
	}
	defer s.Close()
	log.Println("Storage initialized")

	storageService := service.New(s)
	h := transport.NewHandler(storageService)
	mux := http.NewServeMux()
	mux.HandleFunc("/reserve_products", h.ReserveProducts)
	mux.HandleFunc("/release_products", h.ReleaseProducts)
	mux.HandleFunc("/products", h.GetProducts)

	handler := middleware.Logging(mux)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: handler,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server stopped")
			} else {
				log.Fatal("error during server shutdown: ", err)
			}
		}
	}()

	log.Println("Server started")

	<-done
	log.Println("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("stopping server failed: ", err)
	}
}
