package main

import (
	"log"
	"github.com/vfn2002/hdt-response/api"
	"github.com/vfn2002/hdt-response/database"
)

func main() {
	log.Println("Connecting")
	c := database.Connect("mongodb://mongodb:27017")
	log.Println("Connected")

	api.Init(c)
}
