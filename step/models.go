package step

import "github.com/bitrise-io/go-steputils/v2/stepconf"

type Input struct {
	BitriseApiBaseUrl         string          `env:"bitrise_api_base_url,required"`
	BitriseApiAccessToken     stepconf.Secret `env:"bitrise_api_access_token,required"`
	AppSlug                   string          `env:"app_slug,required"`
	AutomaticTestflightUpload bool            `env:"automatic_testflight_upload,opt[true,false]"`
	BundleID                  string          `env:"bundle_id,required"`
	Description               string          `env:"description"`
	ReleaseVersionNumber      string          `env:"release_version_number,required"`
	ReleaseBranch             string          `env:"release_branch"`
	SlackWebhookUrl           string          `env:"slack_webhook_url"`
	TeamsWebhookUrl           string          `env:"teams_webhook_url"`
	Workflow                  string          `env:"workflow"`
	Verbose                   bool            `env:"verbose,opt[true,false]"`
}

type Config struct {
	BitriseApiBaseUrl         string
	BitriseApiAccessToken     string
	AppSlug                   string
	AutomaticTestflightUpload bool
	BundleID                  string
	Description               string
	ReleaseVersionNumber      string
	ReleaseBranch             string
	SlackWebhookUrl           string
	TeamsWebhookUrl           string
	Workflow                  string
}

type Result struct {
	ReleaseUrl  string
	ReleaseSlug string
}
