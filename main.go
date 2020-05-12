package main

import (
	"fmt"

	api "primitivo.fr/applinh/GoFire/apihandler"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
	})	
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.HomeHandler).Methods("GET")
	r.HandleFunc("/{res}",api.GET_ResHandler).Methods("GET")
	r.HandleFunc("/{res}",api.POST_ResHandler).Methods("POST")
	r.HandleFunc("/{res}/{id}",api.PATCH_ResHandler).Methods("PATCH")
   	
	log.Fatal(http.ListenAndServe(":5000", r))
	
	
}