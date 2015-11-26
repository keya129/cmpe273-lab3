package main

import (
	"log"
	"net/http"


)
var HashMap map[int]string
var length float64
var Keylength float64
var KeyMap map[int]string

func main() {
HashMap = make(map[int]string)
KeyMap = make(map[int]string)

HashMap[0]="3000"
HashMap[1]="3001"
HashMap[2]="3002"
length=float64(len(HashMap))
KeyMap[0]=" "
KeyMap[1]="a"
KeyMap[2]="b"
KeyMap[3]="c"
KeyMap[4]="d"
KeyMap[5]="e"
KeyMap[6]="f"
KeyMap[7]="g"
KeyMap[8]="h"
KeyMap[9]="i"
KeyMap[10]="j"
Keylength=float64(len(KeyMap))
KeysAdd()
router := NewRouter()
log.Fatal(http.ListenAndServe(":8081", router))
    

}
