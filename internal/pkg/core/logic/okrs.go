package logic

import (
	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
)

type OkrLogic struct {
	OrgRepository repositories.OrganizationRepository
	OkrRepository repositories.OkrRepository
}

func NewOkrLogic(orgRepository repositories.OrganizationRepository, okrRepository repositories.OkrRepository) OkrLogic {
	return OkrLogic{
		OrgRepository: orgRepository,
		OkrRepository: okrRepository,
	}
}

func (logic *OkrLogic) Validate(newOkr models.Okr) error {

	if newOkr.Title == "" {
		return customerrors.NewArgumentError("title")
	}

	if newOkr.Description == "" {
		return customerrors.NewArgumentError("description")
	}

	if newOkr.OrgId == 0 {
		return customerrors.NewArgumentError("organization")
	}

	if _, err := logic.OrgRepository.Get(newOkr.OrgId); err != nil {
		return customerrors.NewArgumentError("organization")
	}

	return nil
}
