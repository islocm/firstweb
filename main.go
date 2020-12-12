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

// var rNum = regexp.MustCompile(`/d`)              // Has digit(s)
var rAbc = regexp.MustCompile(`selyami`)       // Contains "abc"
var rTarkib = regexp.MustCompile(`tarkib`)     // Contains "abc"
var rChange = regexp.MustCompile(`change`)     // Contains "abc"
var rCom = regexp.MustCompile(`compensation`)  // Contains "abc"
var rComchange = regexp.MustCompile(`comedit`) // Contains "abc"
var rDelete = regexp.MustCompile(`delete`)     // Contains "abc"
var rClean = regexp.MustCompile(`clean`)
var rChop = regexp.MustCompile(`chop`)

func main() {
	e := connection()
	if e != nil {
		fmt.Println(e)
		return
	}
	sessionManager = scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.IdleTimeout = 30 * time.Minute
	mux := http.NewServeMux()
	mux.HandleFunc("/", authBasic(index))
	mux.HandleFunc("/information", authBasic(info))
	mux.HandleFunc("/import", authBasic(element))
	mux.HandleFunc("/Olmazor", authBasic(olmazor))
	mux.HandleFunc("/spisok", adminBasic(spisok))
	mux.HandleFunc("/database", adminBasic(datab))
	mux.HandleFunc("/excel", adminBasic(wrexcel))
	mux.HandleFunc("/execute", adminBasic(hidedb))
	mux.HandleFunc("/islocm", adminBasic(islocm))
	mux.HandleFunc("/otiochsin", authBasic(otiochsin))
	mux.HandleFunc("/importsel", authBasic(selyamiexcel))
	mux.HandleFunc("/zapros", authBasic(zaprost))
	mux.HandleFunc("/rootfile", authBasic(filetofiles))
	mux.HandleFunc("/qwe", authBasic(redspecial))
	mux.HandleFunc("/getall", newexceltofiles)
	mux.HandleFunc("/getselyami", newexceltofilesselyami)
	mux.HandleFunc("/gettarkib", newexceltofilestarkib)
	mux.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./assets"))))
	mux.Handle("/special/", http.StripPrefix("/special", http.FileServer(http.Dir("./files"))))
	http.ListenAndServe(":80", sessionManager.LoadAndSave(mux))
}
