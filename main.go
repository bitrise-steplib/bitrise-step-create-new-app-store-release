package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-steputils/v2/stepenv"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/tmp-bitrise-step-create-new-app-store-release/step"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()

	executor := createReleaseExecutor(logger)
	config, err := executor.ProcessConfig()
	if err != nil {
		logger.Errorf("Process config: %s", err)
		return 1
	}

	result, err := executor.Run(config)
	if err != nil {
		logger.Errorf("Run: %s", err)
		return 1
	}

	logger.Donef("Release successfully created.")
	logger.Println()

	if err := executor.Export(result); err != nil {
		logger.Errorf("Export outputs: %s", err)
		return 1
	}

	return 0
}

func createReleaseExecutor(logger log.Logger) step.ReleaseExecutor {
	envRepository := stepenv.NewRepository(env.NewRepository())
	inputParser := stepconf.NewInputParser(envRepository)

	return step.NewReleaseExecutor(inputParser, envRepository, logger)
}
