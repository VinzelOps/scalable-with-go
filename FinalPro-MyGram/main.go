package main

import (
	"scalable-with-go/FinalPro-MyGram/database"
	"scalable-with-go/FinalPro-MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run()
	if err != nil {
		return
	}
}
