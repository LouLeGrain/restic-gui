package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	method := v["Method"]
	id, _ := strconv.Atoi(v["Id"])

	var status int
	var data interface{}

	switch method {
	case "Snapshots":
		data, err := Snapshots(id)
		if err != nil {
			status = 400
		}
		log.Println(data)
	case "Snapshot":
		data, err := Snapshot(id)
		if err != nil {
			status = 400
		}
		log.Println(data)
	}

	Response := JsonResponse{status, data}

	json.NewEncoder(w).Encode(Response)
}

func Snapshots(id int) (interface{}, error) {
	var data interface{}
	fmt.Println(id)

	return data, nil
}

func Snapshot(id int) (interface{}, error) {
	var data interface{}
	fmt.Println(id)

	return data, nil
}

/*	type post struct {
		Path string
	}
	var p post
	b, _ := ioutil.ReadAll(r.Body)
	//fmt.Print(string(b))

	json.Unmarshal(b, &p)
	exists, _ := utils.CheckFileExists(p.Path)
	w.Header().Set("Content-Type", "application/json")*/
