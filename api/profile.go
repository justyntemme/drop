package api

import (
	"fmt"
	"net/http"

	"gitlab.com/nextwavedevs/drop/dal"

	"github.com/gorilla/mux"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	dal.GetUserProfile(w, r)

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["studioId"])
}
