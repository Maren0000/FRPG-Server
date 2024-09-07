package main

import (
	"FRPGServer/APIHandlers"
	db_commands "FRPGServer/db/commands"
	"context"
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("./.env"); errors.Is(err, os.ErrNotExist) {
		log.Println("No .env file exists")
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	err := db_commands.Opendb(os.Getenv("CUSTOM_SQLITE_PATH"))
	if err != nil {
		log.Println(err)
	}

	Port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /guide.php", APIHandlers.GuideHandler)
	mux.HandleFunc("POST /gw000.php", APIHandlers.Gw000Handler)
	mux.HandleFunc("POST /nativeBridge/native/session.php", APIHandlers.BridgeHandler)

	server := &http.Server{
		Addr:    ":" + Port,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")

	log.Fatal(http.ListenAndServe(":"+Port, mux))
}
