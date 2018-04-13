package main

import (
	"encoding/json"
	"net/http"
)

func SnapShotHandler(w http.ResponseWriter, r *http.Request) {
	/*
		type post struct {
			Path string
		}

		var p post
		b, _ := ioutil.ReadAll(r.Body)
		//fmt.Print(string(b))

		json.Unmarshal(b, &p)
		exists, _ := utils.Exists(p.Path)
		w.Header().Set("Content-Type", "application/json")*/

	response := JsonResponse{200, false}

	json.NewEncoder(w).Encode(response)
}
