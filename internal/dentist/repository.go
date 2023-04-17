package dentist

import (
	"fmt"

	"github.com/oscarpenuela/clinica-odontologica.git/internal/domain"
	"github.com/oscarpenuela/clinica-odontologica.git/pkg/store"
	"github.com/oscarpenuela/clinica-odontologica.git/pkg/web"
)

type IRepository interface {
	GetById(id int) (*domain.Dentist, error)
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) GetById(id int) (*domain.Dentist, error) {
	dent, err := r.Store.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("Dentist_id %d not found", id))
	}
	return dent, nil
}