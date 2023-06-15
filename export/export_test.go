package export

import (
	"errors"
	"github.com/bitrise-steplib/tmp-bitrise-step-create-new-app-store-release/export/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestExport(t *testing.T) {
	key := "key"
	value := "value"

	envRepository := mocks.NewRepository(t)
	envRepository.On("Set", key, value).Return(nil)

	exporter := Exporter{
		EnvRepository: envRepository,
	}

	err := exporter.Export(key, value)
	assert.NoError(t, err)

	envRepository.AssertExpectations(t)
}

func TestExportWithError(t *testing.T) {
	err := errors.New("")

	envRepository := mocks.NewRepository(t)
	envRepository.On("Set", mock.Anything, mock.Anything).Return(err)

	exporter := Exporter{
		EnvRepository: envRepository,
	}

	assert.Error(t, exporter.Export("", ""))

	envRepository.AssertExpectations(t)
}
