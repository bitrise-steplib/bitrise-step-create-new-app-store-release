package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/export"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-steputils/v2/stepenv"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	. "github.com/bitrise-io/go-utils/v2/exitcode"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/step"
)

func main() {
	exitCode := run()
	os.Exit(int(exitCode))
}

func run() ExitCode {
	logger := log.NewLogger()

	executor := createReleaseExecutor(logger)
	config, err := executor.ProcessConfig()
	if err != nil {
		logger.Errorf("Process config: %s", err)
		return Failure
	}

	result, err := executor.Run(config)
	if err != nil {
		logger.Errorf("Run: %s", err)
		return Failure
	}

	if err := executor.Export(result); err != nil {
		logger.Errorf("Export outputs: %s", err)
		return Failure
	}

	return Success
}

func createReleaseExecutor(logger log.Logger) step.ReleaseExecutor {
	envRepository := stepenv.NewRepository(env.NewRepository())
	inputParser := stepconf.NewInputParser(envRepository)
	exporter := export.NewExporter(command.NewFactory(envRepository))

	return step.NewReleaseExecutor(inputParser, envRepository, exporter, logger)
}
