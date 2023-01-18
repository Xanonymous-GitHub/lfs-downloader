package api

import (
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

type httpClientWrapperImpl struct {
	rawClient *req.Client
	logger    *logrus.Logger
}

func CreateHttpClient(logger *logrus.Logger) HttpClient {
	spawnedRawClient := req.C()

	// Enable the debug mode for tracing more details.
	//spawnedRawClient.DevMode()

	return &httpClientWrapperImpl{
		rawClient: spawnedRawClient,
		logger:    logger,
	}
}

func (wrapper *httpClientWrapperImpl) Get(url string) (*req.Response, error) {
	resp, err := wrapper.rawClient.R().
		Get(url)

	if err != nil {
		wrapper.logger.Errorln(err)
		return nil, err
	}

	return resp, nil
}

func (wrapper *httpClientWrapperImpl) Post(url string, body interface{}) (*req.Response, error) {
	resp, err := wrapper.rawClient.R().
		SetBody(body).
		Post(url)

	if err != nil {
		wrapper.logger.Errorln(err)
		return nil, err
	}

	return resp, nil
}
