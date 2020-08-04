package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs"

	_ "github.com/lib/pq"
)

// Name is ...
type Name struct {
	LName []string
}

var sessionManager *scs.SessionManager

func main() {
	e := connection()
	if e != nil {
		fmt.Println(e)
		return
	}
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()
	mux.HandleFunc("/", authBasic(index))
	mux.HandleFunc("/Olmazor", olmazor)
	mux.HandleFunc("/spisok", spisok)
	mux.HandleFunc("/database", datab)
	mux.HandleFunc("/excel", wrexcel)
	mux.HandleFunc("/execdb", hidedb)
	mux.HandleFunc("/otiochsin", otiochsin)
	mux.HandleFunc("/import", importni)
	mux.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":3030", sessionManager.LoadAndSave(mux))
}
