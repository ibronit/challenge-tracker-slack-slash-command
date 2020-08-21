package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibronit/challenge-tracker-slack-slash-command/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomePage).Methods("GET")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		fmt.Print(err)
	}
}
