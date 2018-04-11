package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	type post struct {
		Path string
	}
	var p post
	b, _ := ioutil.ReadAll(r.Body)
	//fmt.Print(string(b))

	json.Unmarshal(b, &p)
	exists, _ := exists(p.Path)
	w.Header().Set("Content-Type", "application/json")

	response := JsonResponse{200, exists}

	json.NewEncoder(w).Encode(response)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
