package logic

import (
	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
)

type OrganizationLogic struct {
	Repository repositories.OrganizationRepository
}

func NewOrganizationLogic(repository repositories.OrganizationRepository) OrganizationLogic {
	return OrganizationLogic{
		Repository: repository,
	}
}

func (logic *OrganizationLogic) Validate(newOrg models.Organization) error {

	if newOrg.Name == "" {
		return customerrors.NewArgumentError("title")
	}

	return nil
}
