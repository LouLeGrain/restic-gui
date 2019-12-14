package main

import "restic-gui/models"

type PageData struct {
	Title   string
	Err     string
	Message string
	Repos   models.Repos
	Backups models.Backups
	Data    string
}

type JsonResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

var PassFile string

var Destination string

type Row struct {
	Id       string `json:"id"`
	DateTime string `json:"date"`
	Host     string `json:"host"`
	Path     string `json:"path"`
}

type RestorePost struct {
	Path 	 string `json:"path"`
	Dest 	 string `json:"dest"`
	File 	 string `json:"file"`
}

type Rows []Row

type Files []string

type Opt map[string]string

type Line string

type Id struct {
	Id interface{}
}
