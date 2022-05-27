/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package fuente

type Repository interface {
	// Lista todas las fuentes
	List() ([]Fuente, error)
	// Lista todas las fuentes por parroquia
	ListByParroquia(parroquia string) ([]Fuente, error)
	// Obtiene una fuente por su identificador
	GetById(id string) (Fuente, error)
}
