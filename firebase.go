package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go"
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
