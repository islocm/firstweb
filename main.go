package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// Name is ...
type Name struct {
	LName []string
}

func main() {
	e := connection()
	if e != nil {
		fmt.Println(e)
		return
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/Olmazor", olmazor)
	http.HandleFunc("/database", datab)
	http.HandleFunc("/excel", wrexcel)
	http.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe("192.168.8.34:3030", nil)
}
