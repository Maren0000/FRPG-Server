package main

import (
	"FRPGServer/APIHandlers"
	db_commands "FRPGServer/db/commands"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	err := db_commands.Opendb()
	if err != nil {
		log.Println(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//IP := os.Getenv("IP_ADDRESS")
	Port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /guide.php", APIHandlers.GuideHandler)
	mux.HandleFunc("POST /gw000.php", APIHandlers.Gw000Handler)
	mux.HandleFunc("POST /nativeBridge/native/session.php", APIHandlers.BridgeHandler)

	log.Fatal(http.ListenAndServe(":"+Port, mux))
}
