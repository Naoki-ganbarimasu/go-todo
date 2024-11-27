package main

import (
	"fmt"
	"go-test-api/db"
	"go-test-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Success Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
