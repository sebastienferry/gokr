package repositories

import (
	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
)

type OrganizationRepository interface {
	GetAll() ([]models.Organization, error)
	Get(id int64) (models.Organization, error)
	Create(models.Organization) (models.Organization, error)
	Update(models.Organization) (models.Organization, error)
	Delete(id int64) error
}

type OrganizationRepositoryInMemory struct {
	Organizations map[int64]models.Organization
}

func NewOrganizationRepositoryInMemory() OrganizationRepository {
	return &OrganizationRepositoryInMemory{
		Organizations: make(map[int64]models.Organization),
	}
}

// Get all organizations.
// TODO: Implement pagination.
func (r *OrganizationRepositoryInMemory) GetAll() ([]models.Organization, error) {
	organizations := make([]models.Organization, 0, len(r.Organizations))
	for _, organization := range r.Organizations {
		organizations = append(organizations, organization)
	}
	return organizations, nil
}

// Get one organization by ID.
func (r *OrganizationRepositoryInMemory) Get(id int64) (models.Organization, error) {
	organization := r.Organizations[id]
	if organization.Id == 0 {
		return models.Organization{}, customerrors.ErrNotFound
	}
	return organization, nil
}

// Creates a new organization.
func (r *OrganizationRepositoryInMemory) Create(organization models.Organization) (models.Organization, error) {
	organization.Id = int64(len(r.Organizations) + 1)
	r.Organizations[organization.Id] = organization
	return organization, nil
}

// Update an existing organization.
func (r *OrganizationRepositoryInMemory) Update(organization models.Organization) (models.Organization, error) {
	organization, err := r.Get(organization.Id)
	if err != nil {
		return models.Organization{}, err
	}
	r.Organizations[organization.Id] = organization
	return organization, nil
}

// Delete an existing organization.
func (r *OrganizationRepositoryInMemory) Delete(id int64) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}
	delete(r.Organizations, id)
	return nil
}
