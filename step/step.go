package step

import (
	"fmt"

	"github.com/bitrise-io/go-steputils/v2/export"
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/api"
)

const (
	releaseUrlKey  = "BITRISE_RELEASE_URL"
	releaseSlugKey = "BITRISE_RELEASE_SLUG"
)

type ReleaseExecutor struct {
	inputParser   stepconf.InputParser
	envRepository env.Repository
	exporter      export.Exporter
	logger        log.Logger
}

func NewReleaseExecutor(inputParser stepconf.InputParser, envRepository env.Repository, exporter export.Exporter, logger log.Logger) ReleaseExecutor {
	return ReleaseExecutor{
		inputParser:   inputParser,
		envRepository: envRepository,
		exporter:      exporter,
		logger:        logger,
	}
}

func (r ReleaseExecutor) ProcessConfig() (Config, error) {
	var input Input
	err := r.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}

	stepconf.Print(input)
	r.logger.Println()
	r.logger.EnableDebugLog(input.Verbose)

	return Config{
		BitriseApiBaseUrl:         input.BitriseApiBaseUrl,
		BitriseApiAccessToken:     input.BitriseApiAccessToken,
		AppSlug:                   input.AppSlug,
		AutomaticTestflightUpload: input.AutomaticTestflightUpload,
		BundleID:                  input.BundleID,
		Description:               input.Description,
		ReleaseVersionNumber:      input.ReleaseVersionNumber,
		ReleaseBranch:             input.ReleaseBranch,
		SlackWebhookUrl:           input.SlackWebhookUrl,
		TeamsWebhookUrl:           input.TeamsWebhookUrl,
		Workflow:                  input.Workflow,
	}, nil
}

func (r ReleaseExecutor) Run(config Config) (Result, error) {
	client := api.NewDefaultAPIClient(config.BitriseApiBaseUrl, config.BitriseApiAccessToken, r.logger)

	parameter := api.CreateReleaseParameter{
		AutomaticTestflightUpload: config.AutomaticTestflightUpload,
		BundleID:                  config.BundleID,
		Description:               config.Description,
		Name:                      config.ReleaseVersionNumber,
		ReleaseBranch:             config.ReleaseBranch,
		SlackWebhookUrl:           config.SlackWebhookUrl,
		TeamsWebhookUrl:           config.TeamsWebhookUrl,
		Workflow:                  config.Workflow,
	}
	response, err := client.CreateRelease(config.AppSlug, parameter)
	if err != nil {
		return Result{}, err
	}

	r.logger.Donef("Release successfully created.")

	return Result{
		ReleaseUrl:  fmt.Sprintf("this-should-be-the-release-url/%s", response.Id),
		ReleaseSlug: response.Id,
	}, nil
}

func (r ReleaseExecutor) Export(result Result) error {
	r.logger.Printf("The following outputs are exported as environment variables:")

	values := map[string]string{
		//TODO: This will be implemented later
		//releaseUrlKey:  result.ReleaseUrl,
		releaseSlugKey: result.ReleaseSlug,
	}

	for key, value := range values {
		err := r.exporter.ExportOutput(key, value)
		if err != nil {
			return err
		}

		r.logger.Donef("$%s = %s", key, value)
	}

	return nil
}
