package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FireStoreClient *firestore.Client
var ClientError error

func InitApp(){
	opt := option.WithCredentialsFile("articleproject-81e10-firebase-adminsdk-h6sj5-5de6d82dab.json")
	config := &firebase.Config{ProjectID: "articleproject-81e10"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	FireStoreClient,ClientError = app.Firestore(context.Background())
	if err != nil{
		log.Fatalln(err)
	}
	if err != nil {
        log.Fatalln(err)
	}

	if ClientError != nil{
		log.Fatalf("error creating client: %v\n", err)
	}
	
}



