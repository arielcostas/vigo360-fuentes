package fuente

import "github.com/jmoiron/sqlx"

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
	if err != nil {
		return []Fuente{}, err
	}
	return fuentes, nil
}

// Lista todas las fuentes por parroquia
func (r *MysqlRepository) ListByParroquia(parroquia string) ([]Fuente, error) {
	var fuentes []Fuente
	err := r.db.Select(&fuentes, `SELECT * FROM fuentes WHERE parroquia = ?`, parroquia)
	if err != nil {
		return []Fuente{}, err
	}
	return fuentes, nil
}

// Obtiene una fuente por su identificador
func (r *MysqlRepository) GetById(id string) (Fuente, error) {
	var fuente Fuente
	err := r.db.QueryRowx("SELECT * FROM fuentes WHERE id=? LIMIT 1", id).StructScan(fuente)
	if err != nil {
		return Fuente{}, err
	}
	return fuente, nil
}
