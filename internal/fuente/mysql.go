/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package fuente

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type MysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(db *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

// Lista todas las fuentes
func (r *MysqlRepository) List() ([]Fuente, error) {
	var fuentes []Fuente
	err := r.db.Select(&fuentes, `SELECT * FROM fuentes`)
	for i, f := range fuentes {
		f.IdImagen = strings.Split(f.Id, "/")[0]
		f.IdImagen = strings.TrimSpace(f.IdImagen)
		fuentes[i] = f
	}
	if err != nil {
		return []Fuente{}, err
	}
	return fuentes, nil
}

// Lista todas las fuentes por parroquia
func (r *MysqlRepository) ListByParroquia(parroquia string) ([]Fuente, error) {
	var fuentes []Fuente
	err := r.db.Select(&fuentes, `SELECT * FROM fuentes WHERE parroquia = ?`, parroquia)
	for i, f := range fuentes {
		f.IdImagen = strings.Split(f.Id, " / ")[0]
		f.IdImagen = strings.TrimSpace(f.IdImagen)
		fuentes[i] = f
	}
	if err != nil {
		return []Fuente{}, err
	}
	return fuentes, nil
}

// Obtiene una fuente por su identificador
func (r *MysqlRepository) GetById(id string) (Fuente, error) {
	var fuente Fuente
	err := r.db.QueryRowx("SELECT * FROM fuentes WHERE id=? LIMIT 1", id).StructScan(&fuente)
	if err != nil {
		return Fuente{}, err
	}
	fuente.IdImagen = strings.Split(fuente.Id, " / ")[0]
	fuente.IdImagen = strings.TrimSpace(fuente.IdImagen)

	return fuente, nil
}
