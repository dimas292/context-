package main

import (
	"defer/route"
	"log"
	"net/http"
)

var (
	APP_PORT = ":4000"
)

func main(){

	route.SetUp()

	log.Printf("server runnng at port %v\n", APP_PORT)
	http.ListenAndServe(APP_PORT, nil)
}

