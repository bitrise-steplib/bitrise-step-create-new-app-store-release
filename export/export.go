package export

import (
	"fmt"

	"github.com/bitrise-io/go-utils/v2/env"
)

type Exporter struct {
	EnvRepository env.Repository
}

func NewExporter(envRepository env.Repository) Exporter {
	return Exporter{
		EnvRepository: envRepository,
	}
}

func (e Exporter) Export(key string, value string) error {
	if err := e.EnvRepository.Set(key, value); err != nil {
		return fmt.Errorf("failed to export output (%s=%s), error: %s", key, value, err)
	}
	return nil
}
