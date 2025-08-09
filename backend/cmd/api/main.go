package main

import (
	"context"
	"encoding/json"
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
		fmt.Fprintf(w, "AnimeConnect backend is running with Firebase! ✅")
	})

	http.HandleFunc("/test-firebase", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		_, _, err := fb.Firestore.Collection("test_collection").Add(ctx, map[string]interface{}{
			"message": "Hello from AnimeConnect!",
		})
		if err != nil {
			http.Error(w, "Failed to write to Firestore: "+err.Error(), http.StatusInternalServerError)
			return
		}

		docs, err := fb.Firestore.Collection("test_collection").Documents(ctx).GetAll()
		if err != nil {
			http.Error(w, "Failed to read from Firestore: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(docs[0].Data())
	})

	fmt.Println("Server is running on http://localhost:" + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
