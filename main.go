package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "cookie",
		Value:    "cookie",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//	DefaultServeMux
	}

	http.HandleFunc("/set_cookie", setCookie)

	server.ListenAndServe()
}
