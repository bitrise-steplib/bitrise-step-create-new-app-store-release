package api

type CreateReleaseParameter struct {
	AutomaticTestflightUpload bool   `json:"automatic_testflight_upload"`
	BundleID                  string `json:"bundle_id"`
	Description               string `json:"description"`
	Name                      string `json:"name"`
	ReleaseBranch             string `json:"release_branch"`
	SlackWebhookUrl           string `json:"slack_webhook_url"`
	TeamsWebhookUrl           string `json:"teams_webhook_url"`
	Workflow                  string `json:"workflow"`
}

type CreateReleaseResponse struct {
	BundleId string `json:"bundle_id"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
}

type errorReponse struct {
	Message string `json:"message"`
}
