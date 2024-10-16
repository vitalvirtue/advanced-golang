package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/vitalvirtue/advanced-golang/Build-A-CRUD-API/routers"
)

func main() {
	r := routers.Router()
	fmt.Println("Starting server on the port 8080...")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}