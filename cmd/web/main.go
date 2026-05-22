package main

import (
	"log"
	"net/http"
)

type application struct {
}	


func main() {
	app := &application{}
	err := http.ListenAndServe(":8000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

