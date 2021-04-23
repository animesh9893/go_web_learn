package main

import (
	firebase "firebase.google.com/go"
	"context"
	"google.golang.org/api/option"
	"log"
)


func addToFirebase(obj User){
	ctx := context.Background()
	sa := option.WithCredentialsFile("../firebaseCred.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
	  log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	  log.Fatalln(err)
	}
	defer client.Close()
	
	_, _, err1 := client.Collection("users").Add(ctx, map[string]interface{}{
	    "Username": obj.Username,
	   	"Password": obj.Password,
	})

	if err1 != nil {
	    log.Fatalf("Failed adding alovelace: %v", err)
	}
}
