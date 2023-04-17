package dentist

import "github.com/oscarpenuela/clinica-odontologica.git/internal/domain"

type IService interface {
	GetById(id int) (*domain.Dentist, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetById(id int)(*domain.Dentist, error) {
	dent , err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dent, nil
}