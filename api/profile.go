package api

import (
	"net/http"
	"fmt"


	"github.com/gorilla/mux"
	"github.com/nextwavedevs/drop/dal"
)


func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	result := dal.getUserProfile(w, r)


	vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["studioId"])
}