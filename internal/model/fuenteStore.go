package model

import (
	"github.com/jmoiron/sqlx"
)

type FuenteStore struct {
	db *sqlx.DB
}

func NewFuenteStore(db *sqlx.DB) FuenteStore {
	return FuenteStore{db: db}
}

func (f FuenteStore) Listar() ([]Fuente, error) {
	var fuentes []Fuente
	err := f.db.Select(&fuentes, `SELECT * FROM fuentes`)
	if err != nil {
		return []Fuente{}, err
	}
	return fuentes, nil
}
