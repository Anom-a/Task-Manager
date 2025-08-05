package main

import (
	"Task-Manager/data"
	"Task-Manager/router"
	"log"
)

func main(){
	data.InitDB()
	r := router.SetupRouter()
	log.Println("Starting server on local port 8080")
	r.Run(":8080")
}