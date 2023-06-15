package step

import (
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/tmp-bitrise-step-create-new-app-store-release/export/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigParsing(t *testing.T) {
	config := Config{
		BitriseApiBaseUrl:         "base-url",
		BitriseApiAccessToken:     "access-token",
		AppSlug:                   "app-slug",
		AutomaticTestflightUpload: false,
		BundleID:                  "bundle-id",
		Description:               "description",
		ReleaseVersionNumber:      "version-number",
		ReleaseBranch:             "branch",
		SlackWebhookUrl:           "slack",
		TeamsWebhookUrl:           "teams",
		Workflow:                  "workflow",
	}

	mockEnvRepository := mocks.NewRepository(t)
	mockEnvRepository.On("Get", "bitrise_api_base_url").Return(config.BitriseApiBaseUrl)
	mockEnvRepository.On("Get", "bitrise_api_access_token").Return(config.BitriseApiAccessToken)
	mockEnvRepository.On("Get", "app_slug").Return(config.AppSlug)

	automaticTestflightUpload := "false"
	if config.AutomaticTestflightUpload {
		automaticTestflightUpload = "true"
	}
	mockEnvRepository.On("Get", "automatic_testflight_upload").Return(automaticTestflightUpload)

	mockEnvRepository.On("Get", "bundle_id").Return(config.BundleID)
	mockEnvRepository.On("Get", "description").Return(config.Description)
	mockEnvRepository.On("Get", "release_version_number").Return(config.ReleaseVersionNumber)
	mockEnvRepository.On("Get", "release_branch").Return(config.ReleaseBranch)
	mockEnvRepository.On("Get", "slack_webhook_url").Return(config.SlackWebhookUrl)
	mockEnvRepository.On("Get", "teams_webhook_url").Return(config.TeamsWebhookUrl)
	mockEnvRepository.On("Get", "workflow").Return(config.Workflow)
	mockEnvRepository.On("Get", "verbose").Return("false")

	inputParser := stepconf.NewInputParser(mockEnvRepository)
	sut := NewReleaseExecutor(inputParser, mockEnvRepository, log.NewLogger())

	receivedConfig, err := sut.ProcessConfig()
	assert.NoError(t, err)
	assert.Equal(t, config, receivedConfig)

	mockEnvRepository.AssertExpectations(t)
}

func TestExport(t *testing.T) {
	result := Result{
		ReleaseUrl:  "release-url",
		ReleaseSlug: "release-slug",
	}

	mockEnvRepository := mocks.NewRepository(t)
	//TODO: Enable it later
	//mockEnvRepository.On("Set", "BITRISE_RELEASE_URL", result.ReleaseUrl).Return(nil)
	mockEnvRepository.On("Set", "BITRISE_RELEASE_SLUG", result.ReleaseSlug).Return(nil)

	inputParser := stepconf.NewInputParser(mockEnvRepository)
	sut := NewReleaseExecutor(inputParser, mockEnvRepository, log.NewLogger())

	err := sut.Export(result)
	assert.NoError(t, err)

	mockEnvRepository.AssertExpectations(t)
}
