package api

import (
	"fmt"
	"net/http"

	"gitlab.com/nextwavedevs/drop/dal"
)

func getProfile(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	uid := vars.Get("uid")

	result := dal.GetUserProfile(w, uid)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "uid: %v\n", vars["uid"])
	fmt.Print(result)
}
