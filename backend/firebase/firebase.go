package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	fb "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/shuakr/AnimeConnect/internal/config"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	Auth      *auth.Client
	Firestore *firestore.Client
}

func InitFirebase(cfg *config.Config) *FirebaseApp {
	ctx := context.Background()

	opt := option.WithCredentialsFile(cfg.FirebaseCredentials)
	app, err := fb.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase App: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth: %v", err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	log.Println("✅ Firebase initialized successfully")

	return &FirebaseApp{
		Auth:      authClient,
		Firestore: firestoreClient,
	}
}
