package fuente

import (
	"errors"
)

type MockRepository struct {
	fail bool
}

func NewMockRepository(fail bool) *MockRepository {
	return &MockRepository{
		fail: fail,
	}
}

/*
	Si MockRepository.fail es true, devuelve un error, si no, devuelve un slice de fuentes vacías
*/
func (r *MockRepository) List() ([]Fuente, error) {
	var (
		fuentes = make([]Fuente, 2)
		err     error
	)

	if r.fail {
		err = errors.New("el test debe fallar")
	}

	return fuentes, err
}

/*
	Si MockRepository.fail es true, devuelve un error.
	Si parroquia='valid', devuelve un slice con dos fuentes vacías
	Si parroquia es cualquier otro valor, devuelve un conjunto vacío
*/
func (r *MockRepository) ListByParroquia(parroquia string) ([]Fuente, error) {
	var (
		fuentes []Fuente
		err     error
	)

	if r.fail {
		err = errors.New("el test debe fallar")
	}
	if parroquia == "valid" {
		fuentes = make([]Fuente, 2)
	}

	return fuentes, err
}

/*
	Si MockRepository.fail es true, devuelve un error, si no, devuelve una fuente vacía
*/
func (r *MockRepository) GetById(id string) (Fuente, error) {
	var err error
	if r.fail {
		err = errors.New("el test debe fallar")
	}
	return Fuente{}, err
}
