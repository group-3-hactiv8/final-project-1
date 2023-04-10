package main

import (
	"final-project-1/database"
	"final-project-1/routers"
)

const port = ":8080"

func main() {
	database.StartDB()
	r := routers.StartApp()

	r.Run(port)
}
