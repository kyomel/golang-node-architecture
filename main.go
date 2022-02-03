package main

import (
	"todolist-golang/src/config"
	"todolist-golang/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	// run all routes
	routes.Routes()
}
