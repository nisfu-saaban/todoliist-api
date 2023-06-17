package main

import (
	"github.com/nisfu-saaban/todoliist-api/app"
	"github.com/nisfu-saaban/todoliist-api/model"
	"github.com/nisfu-saaban/todoliist-api/routers"
)

func main() {
	db := app.ConnectionDatabase()
	model.AutoMigrate(db)
	conn, _ := db.DB()
	defer conn.Close()

	r := routers.SetupRouter()
	r.Run()
}
