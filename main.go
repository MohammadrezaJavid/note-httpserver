package main

import (
	hs "example/httpServers/httpServer"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", hs.MakeHandler(hs.ViewHandler))
	http.HandleFunc("/edit/", hs.MakeHandler(hs.EditHandler))
	http.HandleFunc("/save/", hs.MakeHandler(hs.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
