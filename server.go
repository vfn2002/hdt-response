package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vfn2002/hdt-response/api"
	"github.com/vfn2002/hdt-response/database"
)

// ConnectionString mongo
var ConnectionString string

// Environment env
var Environment string

func main() {
	log.Println(fmt.Sprintf("ENV: %s", os.Getenv("ENV")))
	log.Println(fmt.Sprintf("CONNECTION_STRING: %s", os.Getenv("CONNECTION_STRING")))
	log.Println("Connecting")
	log.Println(ConnectionString)

	c := database.Connect(ConnectionString)

	log.Println("Connected")

	api.Init(c)
}
