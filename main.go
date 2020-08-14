package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/alexedwards/scs"

	_ "github.com/lib/pq"
)

// Name is ...
type Name struct {
	LName []string
}

var sessionManager *scs.SessionManager
var rNum = regexp.MustCompile(`/d`)              // Has digit(s)
var rAbc = regexp.MustCompile(`selyami`)         // Contains "abc"
var rTarkib = regexp.MustCompile(`tarkib`)       // Contains "abc"
var rChange = regexp.MustCompile(`change`)       // Contains "abc"
var rCom = regexp.MustCompile(`compensation`)    // Contains "abc"
var rComchange = regexp.MustCompile(`comchange`) // Contains "abc"

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
	mux.HandleFunc("/information", authBasic(info))
	mux.HandleFunc("/import", authBasic(element))
	mux.HandleFunc("/Olmazor", authBasic(olmazor))
	mux.HandleFunc("/spisok", adminBasic(spisok))
	mux.HandleFunc("/database", adminBasic(datab))
	mux.HandleFunc("/excel", adminBasic(wrexcel))
	mux.HandleFunc("/execdb", adminBasic(hidedb))
	mux.HandleFunc("/islocm", adminBasic(islocm))
	mux.HandleFunc("/otiochsin", authBasic(otiochsin))
	mux.HandleFunc("/importsel", authBasic(selyamiexcel))
	mux.HandleFunc("/zapros", authBasic(zaprost))
	mux.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":3030", sessionManager.LoadAndSave(mux))
}
