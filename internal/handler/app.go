package handler

import (
	"fmt"
	"net/http"
)

type appError struct {
	Error   error
	Log     string
	Mensaje string
	Codigo  int
}

type htmlAppHandler func(w http.ResponseWriter, r *http.Request) *appError

func (fn htmlAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	if ae := fn(w, r); ae != nil {
		fmt.Errorf("%s: %s\n", ae.Log, ae.Error.Error())
		w.WriteHeader(500)
		w.Write([]byte("<!DOCTYPE html>\n<b>Error!!!</b>\n"))
		w.Write([]byte(ae.Mensaje))
	}
}

type jsonAppHandler func(w http.ResponseWriter, r *http.Request) *appError

func (fn jsonAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if ae := fn(w, r); ae != nil {
		fmt.Errorf("%s: %s\n", ae.Log, ae.Error.Error())
		w.WriteHeader(500)
		w.Write([]byte("{error: true, mensaje: \"" + ae.Mensaje + "\"}"))
	}
}
