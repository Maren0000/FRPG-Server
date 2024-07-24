package main

import (
	"FRPGServer/APIHandlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /guide.php", APIHandlers.GuideHandler)
	mux.HandleFunc("POST /gw000.php", APIHandlers.Gw000Handler)
	mux.HandleFunc("POST /nativeBridge/native/session.php", APIHandlers.BridgeHandler)

	log.Fatal(http.ListenAndServe(":80", mux))
}
