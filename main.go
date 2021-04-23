package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	// firebase "firebase.google.com/go"
	// "context"

 //  	"google.golang.org/api/option"
)

// type User struct {
// 	Username string `json:"Username"`
// 	Password string `json:"Password"`
// }

var (
	users []User
)

func homeHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "Home Page")
}

func loginHandler(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-type", "application/json")
	var temp User
	err := json.NewDecoder(req.Body).Decode(&temp)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Error":"Error in Getting info"}`))
		return
	}
	users = append(users,temp)
	fmt.Println(temp)
	addToFirebase(temp)
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(temp)
	res.Write(result)
}

// func addToFirebase(obj User){
// 	ctx := context.Background()
// 	sa := option.WithCredentialsFile("C:/andy/New Folder/go/golang-rest-api-learn-firebase-adminsdk-dsi36-2e932d1da5.json")
// 	app, err := firebase.NewApp(ctx, nil, sa)
// 	if err != nil {
// 	  log.Fatalln(err)
// 	}

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 	  log.Fatalln(err)
// 	}
// 	defer client.Close()
	
// 	_, _, err1 := client.Collection("users").Add(ctx, map[string]interface{}{
// 	    "Username": obj.Username,
// 	   	"Password": obj.Password,
// 	})

// 	if err1 != nil {
// 	    log.Fatalf("Failed adding alovelace: %v", err)
// 	}
// }


func main(){
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/",homeHandler)
	router.HandleFunc("/login",loginHandler).Methods("POST")
	
	log.Println("Server running on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}