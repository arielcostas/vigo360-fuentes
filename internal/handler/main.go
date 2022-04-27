package handler

import (
	"fmt"
	"net/http"

	"vigo360.es/fuentes/internal/database"
	"vigo360.es/fuentes/internal/model"
)

var fuenteStore model.FuenteStore

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fuentes, err := fuenteStore.Listar()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "fuentes: %v\n", fuentes)
}

func Start(port string) {
	fuenteStore = model.NewFuenteStore(database.GetDB())

	http.HandleFunc("/", helloWorld)

	fmt.Printf("Iniciando servidor en %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
