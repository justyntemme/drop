package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/nextwavedevs/drop/dal"
)

func getProfile(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	vars := r.URL.Query()
	uid := vars.Get("uid")

	result := dal.GetUserById(ctx, w, uid)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v\n", result)

}
