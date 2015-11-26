package main

import (
	"log"
	"net/http"

)
var Map map[int]HashMap
var index int
func main() {
Map = make(map[int]HashMap)
index=0
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3001", router))
    

}
