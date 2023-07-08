package main

import (
	"example/httpServers/httpServer"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", httpServer.ViewHandler)
	http.HandleFunc("/edit/", httpServer.EditHandler)
	http.HandleFunc("/save/", httpServer.SaveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
