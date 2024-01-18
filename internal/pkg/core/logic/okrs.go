package logic

import (
	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
)

type OkrLogic struct {
	Repository repositories.OkrRepository
}

func NewOkrLogic(repository repositories.OkrRepository) OkrLogic {
	return OkrLogic{
		Repository: repository,
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

	return nil
}
