package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/retry"
	"github.com/bitrise-io/go-utils/v2/log"
)

const (
	timeout = 30
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type DefaultAPIClient struct {
	httpClient HttpClient
	authToken  stepconf.Secret
	baseURL    string
	logger     log.Logger
}

func NewDefaultAPIClient(baseURL string, authToken stepconf.Secret, logger log.Logger) DefaultAPIClient {
	httpClient := retry.NewHTTPClient().StandardClient()
	httpClient.Timeout = time.Second * timeout

	return DefaultAPIClient{
		httpClient: httpClient,
		authToken:  authToken,
		baseURL:    baseURL,
		logger:     logger,
	}
}

func (c DefaultAPIClient) CreateRelease(appSlug string, params CreateReleaseParameter) (CreateReleaseResponse, error) {
	req, err := c.request(appSlug, params)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		c.logger.Warnf("request dump failed: %s", err)
	} else {
		c.logger.Debugf("Request dump: %s", string(dump))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.Printf(" [!] Failed to close response body: %+v", err)
		}
	}()

	dump, err = httputil.DumpResponse(resp, true)
	if err != nil {
		c.logger.Warnf("response dump failed: %s", err)
	} else {
		c.logger.Debugf("Response dump: %s\n", string(dump))
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		errResponse, err := c.parseError(resp)
		if err != nil {
			c.logger.Warnf("Failed to parse failure reason from the response: %s", err)
			return CreateReleaseResponse{}, fmt.Errorf("request to %s has status code %d (should be 2XX)", req.URL.String(), resp.StatusCode)
		} else {
			return CreateReleaseResponse{}, fmt.Errorf("request to %s has status code %d (should be 2XX): %s", req.URL.String(), resp.StatusCode, errResponse.Message)
		}
	}

	parsedResp, err := c.parseModel(resp)
	if err != nil {
		return CreateReleaseResponse{}, fmt.Errorf("release successfully created but response couldn't be parsed: %s", err)
	}
	return parsedResp, nil
}

func (c DefaultAPIClient) request(appSlug string, params CreateReleaseParameter) (*http.Request, error) {
	requestPath := fmt.Sprintf("%s/v0.1/apps/%s/releases/app-store", c.baseURL, appSlug)

	paramBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, requestPath, bytes.NewBuffer(paramBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", string(c.authToken))

	return req, nil
}

func (c DefaultAPIClient) parseError(resp *http.Response) (errorReponse, error) {
	var body []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorReponse{}, err
	}

	var response errorReponse
	if err := json.Unmarshal(body, &response); err != nil {
		return errorReponse{}, err
	}

	return response, nil
}

func (c DefaultAPIClient) parseModel(resp *http.Response) (CreateReleaseResponse, error) {
	var body []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateReleaseResponse{}, err
	}

	var response CreateReleaseResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return CreateReleaseResponse{}, err
	}

	return response, nil
}
