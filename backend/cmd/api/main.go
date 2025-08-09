package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shuakr/AnimeConnect/firebase"
	"github.com/shuakr/AnimeConnect/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fb := firebase.InitFirebase(cfg)
	defer fb.Firestore.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AnimeConnect backend is running with Firebase!")
	})

	fmt.Println("Server is running on http://localhost:" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
