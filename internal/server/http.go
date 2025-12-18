package server

import (
	"log"
	"net/http"

	"example.com/go-gateway/internal/middleware"
	"example.com/go-gateway/internal/proxy"
)

func Start() {
	mux := http.NewServerMux()

	handler := middleware.Chain(
		proxy.ReverseProxy(),
		middleware.Logging,
		middleware.Policy,
	)

	mux.Handle("/api", handler)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServer(":8080", mux))
}
