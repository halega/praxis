package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
}
