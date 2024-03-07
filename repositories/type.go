package repositories

import (
	"database/sql"

	"github.com/Elsharaky/Cars-API.git/config"
	"github.com/Elsharaky/Cars-API.git/models"
)

type TypeRepository interface {
	CreateType(t *models.Type) error
	GetTypes() ([]*models.Type, error)
}

type typeRepository struct {
	db *sql.DB
}

func NewTypeRepository() *typeRepository {
	return &typeRepository{
		db: config.GetDB(),
	}
}

func (tr *typeRepository) CreateType(t *models.Type) error {
	query := `INSERT INTO types (name) VALUES($1) RETURNING id, name`
	return tr.db.QueryRow(query, t.Name).Scan(&t.ID, &t.Name)
}

func (tr *typeRepository) GetTypes() ([]*models.Type, error) {
	query := `SELECT id, name FROM types`
	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	types := []*models.Type{}
	for rows.Next() {
		var t models.Type
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		types = append(types, &t)
	}

	return types, nil
}
