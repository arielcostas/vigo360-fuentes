/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package fuente

type Fuente struct {
	Id       string
	IdImagen string
	Titulo   string

	Ubicacion string
	Parroquia string
	Latitud   string
	Longitud  string

	Caudal    string
	Calidad   string
	Origen    string
	Lavadero  bool
	Analizada string
	Nota      string
}
