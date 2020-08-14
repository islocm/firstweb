package main

import "net/http"

type handler func(w http.ResponseWriter, r *http.Request)

func authBasic(handler handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != "username" || password != "12345" {
			http.Error(w, "Not authorized", 401)
			return
		}

		handler(w, r)
	}
}

func adminBasic(handler handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != "islocm" || password != "12311231345345" {
			http.Error(w, "Not authorized", 401)
			return
		}

		handler(w, r)
	}
}
