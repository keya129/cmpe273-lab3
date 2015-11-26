package main

import (
	"encoding/json"
	"fmt"
    "net/http"
    "strconv"
    "io/ioutil"
    "github.com/gorilla/mux"	
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func KeyShow(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
        keyId := vars["key_id"]
		l, err := strconv.Atoi(keyId)
		if err != nil {
			panic(err)
		}
        var newMap HashMap
        for i:=0;i<len(Map);i++{
        temp:=Map[i].Key
        if(l==temp){
        newMap.Key=Map[i].Key
        newMap.Value=Map[i].Value
        }
        }
        res2B, _ :=json.Marshal(newMap)
        w.Header().Set("Content-Type", "application/json;")
        w.WriteHeader(200)
        fmt.Fprintf(w,"%s",res2B)

        

}
func KeyAdd(w http.ResponseWriter, r *http.Request) {
var dat HashMap
	body, _ := ioutil.ReadAll(r.Body)
    err := json.Unmarshal(body, &dat);
    if err != nil {
  	panic(err)
    }
    Map[index]=dat
    fmt.Println(index)
    index++
}
func KeysShow(w http.ResponseWriter, r *http.Request) {
value := make([]HashMap, len(Map))

    for i:=0;i<len(Map);i++{
    value[i]=Map[i]

    }
res2B, _ :=json.Marshal(value)
        w.Header().Set("Content-Type", "application/json;")
        w.WriteHeader(200)
        fmt.Fprintf(w,"%s",res2B)
}
func KeyUpdate(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
    keyId := vars["key_id"]
		l, err := strconv.Atoi(keyId)
		if err != nil {
			panic(err)
		}
    value := vars["value"]
    var newMap HashMap
    for i:=0;i<len(Map);i++{
    temp:=Map[i].Key
    if(l==temp){
    newMap.Key=Map[i].Key
    newMap.Value=value
    Map[i]=newMap
            }
     }
    res2B, _ :=json.Marshal(newMap)
    w.Header().Set("Content-Type", "application/json;")
    w.WriteHeader(200)
    fmt.Fprintf(w,"%s",res2B)

}
