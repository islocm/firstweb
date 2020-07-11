package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tem.Execute(w, nil)
}

func olmazor(w http.ResponseWriter, r *http.Request) {
	x := r.FormValue("name")
	y := r.FormValue("qwe")
	fmt.Println(x)
	fmt.Println(y)
	qwe, err := template.ParseFiles("template/Olmazor.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	qwe.Execute(w, nil)

}
