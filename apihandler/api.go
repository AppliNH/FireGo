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
	vars := mux.Vars(r)
	res := vars["res"]

	result := db.ReadRes(res)

	if val, ok := vars["id"]; ok {

		if _, ok := result[val]; ok {
			json.NewEncoder(w).Encode(result[val])
		} else {
			json.NewEncoder(w).Encode(map[string]string{"statusCode": "400", "success": "false", "error": "item not found"})
		}
		
	} else {
		json.NewEncoder(w).Encode(result)
	}
	
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

	uuid := db.WriteItem(res, data)
	response := map[string]string{"statusCode": "200", "success": "true", "id": uuid, "res": res}
	json.NewEncoder(w).Encode(response)
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
	response := map[string]string{"statusCode": "200", "success": "true", "id": id, "res": res}
	json.NewEncoder(w).Encode(response)
}
