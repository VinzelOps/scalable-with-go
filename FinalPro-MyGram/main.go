package main

import (
	"github.com/VinzelOps/scalable-with-go/FinalPro-MyGram/database"
	"github.com/VinzelOps/scalable-with-go/FinalPro-MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run()
	if err != nil {
		return
	}
}
