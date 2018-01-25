package main

import (
	"dsound/db"
	"dsound/routes"
)

func main() {
	db.EnsureIndices()
	routes.StartServer()
}
