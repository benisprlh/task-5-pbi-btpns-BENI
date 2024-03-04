package main

import (
	"project/database"
	"project/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
