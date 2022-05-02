package fuente

type Repository interface {
	// Lista todas las fuentes
	List() ([]Fuente, error)
	// Lista todas las fuentes por parroquia
	ListByParroquia(parroquia string) ([]Fuente, error)
	// Obtiene una fuente por su identificador
	GetById(id string) (Fuente, error)
}
