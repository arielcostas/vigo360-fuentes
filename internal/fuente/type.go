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
