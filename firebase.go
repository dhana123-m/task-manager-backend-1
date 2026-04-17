package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	key := os.Getenv("FIREBASE_CREDENTIALS")

	opt := option.WithCredentialsJSON([]byte(key))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	App = app
}
