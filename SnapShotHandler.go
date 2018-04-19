package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SnapShotHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Println(vars)

	route := mux.Route{}

	fmt.Println(route)

	/*
		type post struct {
			Path string
		}

		var p post
		b, _ := ioutil.ReadAll(r.Body)
		//fmt.Print(string(b))

		json.Unmarshal(b, &p)
		exists, _ := utils.CheckFileExists(p.Path)
		w.Header().Set("Content-Type", "application/json")*/

	response := JsonResponse{200, false}

	json.NewEncoder(w).Encode(response)
}
