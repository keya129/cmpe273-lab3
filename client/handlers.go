package main

import (
	"encoding/json"
	"fmt"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"	
    "math"
    "log"
    "io/ioutil"
	"bytes"

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
        x:=math.Mod(float64(l),length)
        port:=HashMap[int(x)]
        url := "http://localhost:"+port+"/keys/"+keyId
                fmt.Println(url)

	  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  	log.Fatal(err)
  }
  req.SetBasicAuth("<token>", "x-oauth-basic")

  client := http.Client{}
  res, err := client.Do(req)
  if err != nil {
  	log.Fatal(err)
  }

  log.Println("StatusCode:", res.StatusCode)
  var dat HashVal
  // read body
  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
  	log.Fatal(err)
  }

  err = json.Unmarshal(body, &dat);if err != nil {
  panic(err)
  }
  var newdat HashVal
  newdat.Key=dat.Key
  newdat.Value=dat.Value
res2B, _ :=json.Marshal(newdat)
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
    var newMap HashVal
    newMap.Key=l
    newMap.Value=value
    res2B, _ :=json.Marshal(newMap)
    x:=math.Mod(float64(l),length)
    port:=HashMap[int(x)]
    url := "http://localhost:"+port+"/keys/"+keyId+"/"+value
    fmt.Println(url)
	var jsonStr = []byte(res2B)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
    w.Header().Set("Content-Type", "application/json;")
    w.WriteHeader(200)
    //fmt.Fprintf(w,"%s",res21B)
    fmt.Fprintf(w,"%s","200")


}
func KeysAdd(){
for j:=1;j<len(KeyMap);j++{
x:=math.Mod(float64(j),length)
port:=HashMap[int(x)]
url := "http://localhost:"+port+"/key"
fmt.Println(url)
var newMap HashVal
newMap.Key=j
newMap.Value=KeyMap[j]
fmt.Println(newMap)
res2B, _ :=json.Marshal(newMap)   
var jsonStr = []byte(res2B)
req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
req.Header.Set("Content-Type", "application/json")
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
panic(err)
}	
fmt.Println(resp)

}
}