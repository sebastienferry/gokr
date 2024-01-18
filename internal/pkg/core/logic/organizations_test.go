package logic

import (
	"testing"

	customerrors "github.com/sebastienferry/gokr/core/errors"
	"github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
	"github.com/stretchr/testify/assert"
)

var (
	orgaRepo  = repositories.NewOrganizationRepositoryInMemory()
	orgaLogic = NewOrganizationLogic(orgaRepo)
)

func TestOrganizationLogic_ValidateTitle(t *testing.T) {
	// Arrange & Act
	okr := models.Okr{}
	err := okrLogic.Validate(okr)

	// Assert
	if assert.Error(t, err) {
		assert.Equal(t, customerrors.NewArgumentError("title"), err)
	}
}

func TestOrganizationLogic_ValidateDescription(t *testing.T) {
	// Arrange & Act
	okr := models.Okr{Title: "title"}
	err := okrLogic.Validate(okr)

	// Assert
	if assert.Error(t, err) {
		assert.Equal(t, customerrors.NewArgumentError("description"), err)
	}
}

func TestOrganizationLogic_ValidateOkr(t *testing.T) {
	// Arrange & Act
	okr := models.Okr{Title: "title", Description: "description", OrgId: 1}
	err := okrLogic.Validate(okr)

	// Assert
	assert.Nil(t, err)
}
