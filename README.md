# Prepare App Store Release

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-create-new-app-store-release?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/releases)

The Step creates an App Store release with the Release Management feature.

<details>
<summary>Description</summary>

Create and configure a new release to App Store Connect with Release Management.
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `bundle_id` | The bundle ID of the app to be released.   Release Management requires the bundle ID of the iOS application being released to the App Store Connect. | required |  |
| `release_version_number` | The version number of the app you are releasing.  Numbering should follow software versioning conventions (1.0, 1.0.0). | required |  |
| `automatic_testflight_upload` | Indicates whether or not to upload every release candidate build automatically to TestFlight.  Release Management will deploy each release candidate to TestFlight automatically if this setting is enabled. Note: This feature requires the release branch and Workflow to be set. | required | `false` |
| `description` | An internal description of the release, it won't be propagated to the App Store.  This description will not be visible on the App Store Connect or available for the end user. |  |  |
| `release_branch` | The branch you created for this version of your app.  This branch is called mostly: release-1.0, release-october, main, etc. |  |  |
| `workflow` | The workflow that generates your an .xcarchive or an App Store signed .ipa artifact.  Make sure that the Workflow generates the artifact for the same Bundle Identifier you provided for this Step as a step input. Release Management will ignore any other .xcarchive App Store signed .ipa with different bundle ID. |  |  |
| `slack_webhook_url` | The Slack webhook URL to use for sending Slack notifications.  By providing a Slack webhook URL, Release Management will send automatic messages for the following events: ``` ┌─────────────────────┬───────────────────────────────────────┐ │ Stage               │ Event                                 │ ├─────────────────────┼───────────────────────────────────────┤ │ Release candidate   │ Release candidate changed             │ │ TestFlight upload   │ Upload and processing finished        │ │ Approvals           │ Release approved                      │ │ App Store review    │ Release sent for review               │ │                     │ Status of review submission changed   │ │ Release             │ Release started                       │ │                     │ Release finished                      │ └─────────────────────┴───────────────────────────────────────┘ ``` For more information go to our [devcenter notification page.](https://devcenter.bitrise.io/en/release-management/enabling-slack-notifications-for-release-management-events.html) |  |  |
| `teams_webhook_url` | The Teams webhook URL to use for sending Teams notifications.  By providing a Teams webhook URL, Release Management will send automatic messages for the following events: ``` ┌─────────────────────┬───────────────────────────────────────┐ │ Stage               │ Event                                 │ ├─────────────────────┼───────────────────────────────────────┤ │ Release candidate   │ Release candidate changed             │ │ TestFlight upload   │ Upload and processing finished        │ │ Approvals           │ Release approved                      │ │ App Store review    │ Release sent for review               │ │                     │ Status of review submission changed   │ │ Release             │ Release started                       │ │                     │ Release finished                      │ └─────────────────────┴───────────────────────────────────────┘ ``` For more information go to our [devcenter notification page.](https://devcenter.bitrise.io/en/release-management/enabling-slack-notifications-for-release-management-events.html) |  |  |
| `bitrise_api_access_token` | Your access token.  To acquire a `Personal Access Token` for your user, sign in with that user on [bitrise.io](https://bitrise.io), go to your `Account Settings` page, and select the [Security tab](https://www.bitrise.io/me/profile#/security) on the left side. | required, sensitive |  |
| `bitrise_api_base_url` | The base URL of the Bitrise API used to process the download requests.  By default the step will use the official Bitrise Public API, you don’t need to change this setting. | required | `https://api.bitrise.io` |
| `app_slug` | The identifier of the Bitrise app for which to create a new release.  By default, the Step will create a new release for the same Bitrise App. | required | `$BITRISE_APP_SLUG` |
| `verbose` | Enable logging additional information for debugging. | required | `false` |
</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `BITRISE_RELEASE_SLUG` | Unique identifier of the newly created release. |
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Note: this step's end-to-end tests (defined in e2e/bitrise.yml) are working with secrets which are intentionally not stored in this repo. External contributors won't be able to run those tests. Don't worry, if you open a PR with your contribution, we will help with running tests and make sure that they pass.

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
