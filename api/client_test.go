package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-create-new-app-store-release/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	authToken = "auth-token"
	baseURL   = "some-base-url"
	appSlug   = "app-slug"
)

func TestBaseParametersAreCorrect(t *testing.T) {
	// Given
	apiClient, mockHttpClient := createSutAndMock(t)

	var request http.Request
	setupMockNetworking(mockHttpClient, &request, "", 0)

	// When
	_, _ = apiClient.CreateRelease(appSlug, CreateReleaseParameter{})

	// Then
	assert.Equal(t, fmt.Sprintf("%s/v0.1/apps/%s/releases/app-store", baseURL, appSlug), request.URL.String())
	assert.Equal(t, authToken, request.Header.Get("Authorization"))

	mockHttpClient.AssertExpectations(t)
}

func TestParametersHandledCorrectly(t *testing.T) {
	// Given
	apiClient, mockHttpClient := createSutAndMock(t)

	var request http.Request
	setupMockNetworking(mockHttpClient, &request, "", 0)

	params := CreateReleaseParameter{
		AutomaticTestflightUpload: true,
		BundleID:                  "1",
		Description:               "2",
		Name:                      "3",
		ReleaseBranch:             "4",
		SlackWebhookUrl:           "5",
		TeamsWebhookUrl:           "6",
		Workflow:                  "7",
	}

	// When
	_, _ = apiClient.CreateRelease(appSlug, params)

	// Then
	var received CreateReleaseParameter
	err := json.NewDecoder(request.Body).Decode(&received)
	assert.NoError(t, err)
	assert.Equal(t, params, received)

	mockHttpClient.AssertExpectations(t)
}

func TestHandlesResponse(t *testing.T) {
	tests := []struct {
		name                 string
		responseStatusCode   int
		responseBody         string
		wantsError           bool
		expectedErrorMessage string
		expectedOutput       CreateReleaseResponse
	}{
		{
			name:                 "Failure 1",
			responseStatusCode:   199,
			responseBody:         "{\"message\": \"Something went wrong.\"}",
			wantsError:           true,
			expectedErrorMessage: "Something went wrong.",
		},
		{
			name:                 "Failure 2",
			responseStatusCode:   301,
			responseBody:         "{\"message\": \"There was an error.\"}",
			wantsError:           true,
			expectedErrorMessage: "There was an error.",
		},
		{
			name:               "Successful response",
			responseStatusCode: 200,
			responseBody: `{
"bundle_id": "bundle-id",
"id": "id",
"name": "name",
"platform": "platform",
"status": "status"
}`,
			wantsError: false,
			expectedOutput: CreateReleaseResponse{
				BundleId: "bundle-id",
				Id:       "id",
				Name:     "name",
				Platform: "platform",
				Status:   "status",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			apiClient, mockHttpClient := createSutAndMock(t)

			setupMockNetworking(mockHttpClient, nil, tt.responseBody, tt.responseStatusCode)

			// When
			data, err := apiClient.CreateRelease(appSlug, CreateReleaseParameter{})

			// Then
			if tt.wantsError {
				assert.True(t, strings.Contains(err.Error(), tt.expectedErrorMessage))
			} else {
				assert.Equal(t, tt.expectedOutput, data)
			}

			mockHttpClient.AssertExpectations(t)
		})
	}
}

func createSutAndMock(t *testing.T) (DefaultAPIClient, *mocks.HttpClient) {
	mockHttpClient := mocks.NewHttpClient(t)
	apiClient := DefaultAPIClient{
		httpClient: mockHttpClient,
		authToken:  authToken,
		baseURL:    baseURL,
		logger:     log.NewLogger(),
	}

	return apiClient, mockHttpClient
}

func setupMockNetworking(mockHttpClient *mocks.HttpClient, request *http.Request, body string, statusCode int) {
	response := http.Response{Body: io.NopCloser(bytes.NewReader([]byte(body)))}
	response.StatusCode = statusCode

	mockHttpClient.On("Do", mock.Anything).Return(&response, nil).Run(func(args mock.Arguments) {
		if request == nil {
			return
		}
		value := args.Get(0).(*http.Request)
		*request = *value
	})
}
