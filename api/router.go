package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/nextwavedevs/drop/dal"
)

func StartApi() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/create/user", CreateUserHandler)

	//get
	r.HandleFunc("/get/studio", GetStudioHandler)
	r.HandleFunc("/get/user", GetProfileHandler)

	log.Fatal(srv.ListenAndServe())
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Create Profile Endpoint Hit")
	dal.CreateProfile(w, r)

}

func GetStudioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["studioId"])
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	getProfile(w, r)
}
