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

type appHandler func(w http.ResponseWriter, r *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ae := fn(w, r); ae != nil {
		fmt.Errorf("%s: %s\n", ae.Log, ae.Error.Error())
		w.WriteHeader(500)
		w.Write([]byte("<!DOCTYPE html>\n<b>Error!!!</b>\n"))
		w.Write([]byte(ae.Mensaje))
	}
}
