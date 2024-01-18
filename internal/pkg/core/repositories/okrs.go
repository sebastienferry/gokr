package repositories

import (
	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
)

type OkrRepository interface {
	GetAll() ([]models.Okr, error)
	Get(id int64) (models.Okr, error)
	Create(models.Okr) (models.Okr, error)
	Update(models.Okr) (models.Okr, error)
	Delete(id int64) error
}

type OkrRepositoryInMemory struct {
	Okrs map[int64]models.Okr
}

// Get all OKRs.
// TODO: Implement pagination.
func (r *OkrRepositoryInMemory) GetAll() ([]models.Okr, error) {
	okrs := make([]models.Okr, 0, len(r.Okrs))
	for _, okr := range r.Okrs {
		okrs = append(okrs, okr)
	}
	return okrs, nil
}

// Get one OKR by ID.
// Returns an error if the OKR is not found.
func (r *OkrRepositoryInMemory) Get(id int64) (models.Okr, error) {
	okr := r.Okrs[id]
	if okr.Id == 0 {
		return models.Okr{}, customerrors.ErrNotFound
	}
	return okr, nil
}

// Creates a new OKR.
func (r *OkrRepositoryInMemory) Create(okr models.Okr) (models.Okr, error) {
	okr.Id = int64(len(r.Okrs) + 1)
	r.Okrs[okr.Id] = okr
	return okr, nil
}

// Update an existing OKR.
func (r *OkrRepositoryInMemory) Update(okr models.Okr) (models.Okr, error) {
	okr, err := r.Get(okr.Id)
	if err != nil {
		return models.Okr{}, err
	}
	r.Okrs[okr.Id] = okr
	return okr, nil
}

// Delete an existing OKR.
func (r *OkrRepositoryInMemory) Delete(id int64) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}
	delete(r.Okrs, id)
	return nil
}

// NewOkrRepositoryInMemory creates a new in-memory OKR repository.
func NewOkrRepositoryInMemory() OkrRepository {
	return &OkrRepositoryInMemory{
		Okrs: make(map[int64]models.Okr),
	}
}
