package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vigo360.es/fuentes/internal/database"
	"vigo360.es/fuentes/internal/model"
)

var fuenteStore model.FuenteStore

func helloWorld(w http.ResponseWriter, r *http.Request) *appError {
	fuentes, err := fuenteStore.Listar()
	if err != nil {
		return &appError{err, "error obtaining fuentes", "Error obteniendo datos", 500}
	}
	fmt.Fprintf(w, "fuentes: %v\n", fuentes)
	return nil
}

func helloJson(w http.ResponseWriter, r *http.Request) *appError {
	fuentes, err := fuenteStore.Listar()
	if err != nil {
		return &appError{err, "error obtaining fuentes", "Error obteniendo datos", 500}
	}
	salida, err := json.MarshalIndent(fuentes, "", "\t")
	if err != nil {
		return &appError{err, "error encoding JSON", "Error procesando datos", 500}
	}
	w.Write(salida)
	return nil
}

func Start(port string) {
	fuenteStore = model.NewFuenteStore(database.GetDB())

	http.Handle("/api", jsonAppHandler(helloJson))
	http.Handle("/", htmlAppHandler(helloWorld))

	fmt.Printf("Iniciando servidor en %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
