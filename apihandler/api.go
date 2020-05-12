package apihandler


import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	db "primitivo.fr/applinh/GoFire/db"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
}

func GET_ResHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	res := mux.Vars(r)["res"]
	fmt.Println(res)
	result := db.ReadRes(res)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func POST_ResHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := mux.Vars(r)["res"]
	
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)
	
    if err != nil {
		fmt.Println(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	db.WriteItem(res,data)
	fmt.Println(res)
	fmt.Println(data)
}
func PATCH_ResHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := mux.Vars(r)["res"]
	id := mux.Vars(r)["id"]
	
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	
	db.UpdateItem(res, id, data)
}