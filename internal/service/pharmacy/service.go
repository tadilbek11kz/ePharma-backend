package pharmacy

import (
	"context"

	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
)

//go:generate mockery --name Service
type Service interface {
	CreatePharmacy(ctx context.Context, req model.CreatePharmacyRequest) (pharmacy model.Pharmacy, err error)
	GetAllPharmacies(ctx context.Context) (pharmacies []model.Pharmacy, err error)
	GetPharmacy(ctx context.Context, id string) (pharmacy model.Pharmacy, err error)
	UpdatePharmacy(ctx context.Context, id string, req model.UpdatePharmacyRequest) (pharmacy model.Pharmacy, err error)
	DeletePharmacy(ctx context.Context, id string) (err error)
}

type service struct {
	st *store.RepositoryStore
}

func New(st *store.RepositoryStore) (srv Service) {
	srv = &service{
		st: st,
	}
	srv = WithLogging(srv)
	return
}

func (s *service) CreatePharmacy(ctx context.Context, req model.CreatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	return s.st.PharmacyRepository.CreatePharmacy(req)
}

func (s *service) GetAllPharmacies(ctx context.Context) (pharmacies []model.Pharmacy, err error) {
	return s.st.PharmacyRepository.GetAllPharmacies()
}

func (s *service) GetPharmacy(ctx context.Context, id string) (pharmacy model.Pharmacy, err error) {
	return s.st.PharmacyRepository.GetPharmacy(id)
}

func (s *service) UpdatePharmacy(ctx context.Context, id string, req model.UpdatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	return s.st.PharmacyRepository.UpdatePharmacy(id, req)
}

func (s *service) DeletePharmacy(ctx context.Context, id string) (err error) {
	return s.st.PharmacyRepository.DeletePharmacy(id)
}
