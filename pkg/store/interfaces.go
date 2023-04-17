package store

import "github.com/oscarpenuela/clinica-odontologica.git/internal/domain"

type StoreInterface interface {
	Read(id int) (*domain.Dentist, error)
}