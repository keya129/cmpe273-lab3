package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"KeyShow",
		"GET",
		"/keys/{key_id}",
		KeyShow,
	},
	Route{
		"KeyUpdate",
		"PUT",
		"/keys/{key_id}/{value}",
		KeyUpdate,
	},
    
}
