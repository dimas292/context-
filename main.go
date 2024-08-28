package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index called ...")
	time := time.Now()
	w.Write([]byte(fmt.Sprint(time.Hour(), ":", time.Minute())))
}

func handleFunc() {
	http.HandleFunc("/", index)
}

func main() {

	port := ":9000"
	handleFunc()
	startServer(port)
}

func startServer(port string) {
	log.Printf("server running at port %v\n", port)
	http.ListenAndServe(port, nil)
}
