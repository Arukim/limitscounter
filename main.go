package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var c *cache.Cache

func init() {
	c = cache.New(1*time.Minute, 30*time.Second)
}

func checkLimit(key string) bool {
	node, found := c.Get(key)
	if !found {
	}
	return true
}

func limitsHandler(w http.ResponseWriter, r *http.Request) {
	resp := checkLimit(mux.Vars(r)["key"])	

	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/limitscounter/{key}", limitsHandler)
	http.Handle("/", r)
	http.ListenAndServe(":666",nil)
}
