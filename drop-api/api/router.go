// package api

// import (
// 	"net/http"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// func StartApi() {
// 	r := mux.NewRouter()
// 	srv := &http.Server{
//         Handler:      r,
//         Addr:         "127.0.0.1:8080",
//         // Good practice: enforce timeouts for servers you create!
//         WriteTimeout: 15 * time.Second,
//         ReadTimeout:  15 * time.Second,
// 	}
	
// 	//Routes
// 	r.HandleFunc("/get/studio", GetStudioHandler)
// 	r.HandleFunc("/get/user", GetProfileHandler)
//     log.Fatal(srv.ListenAndServe())
// }

// func GetStudioHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
//     w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Category: %v\n", vars["studioId"])
// }