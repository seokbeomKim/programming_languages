package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	// var err error
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	converted, err := strconv.ParseFloat(price, 32)
	if err != nil {
		log.Fatal("Failed convert value to dollar\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db[item] = dollars(converted)
	fmt.Fprintf(w, "Updated: %s: %s\n", item, price)
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	// http 편의상 DefaultServeMux를 제공하므로 Mux 인스턴르로 nil을
	// 전달하면 기본값을 사용하게 된다.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
