package store

import (
	"database/sql"

	"github.com/oscarpenuela/clinica-odontologica.git/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read(id int)(*domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentist.Matricula, &dentist.Apellido, &dentist.Nombre)
	if err != nil {
		return nil, err
	}

	return &dentist, nil
}