package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shuakr/AnimeConnect.git/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AnimeConnect backend is running on port %s", cfg.Port)

	})
	fmt.Println("Server is running on http://localhost:" + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
