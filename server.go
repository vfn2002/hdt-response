package main

import (
	"github.com/vfn2002/hdt-response/api"
	"github.com/vfn2002/hdt-response/database"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	api.Init(db)
}
