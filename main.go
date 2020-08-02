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

	http.HandleFunc("/", authBasic(index))
	http.HandleFunc("/Olmazor", olmazor)
	http.HandleFunc("/database", datab)
	http.HandleFunc("/excel", wrexcel)
	http.HandleFunc("/execdb", hidedb)
	http.HandleFunc("/otiochsin", otiochsin)
	http.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":3030", nil)
}
