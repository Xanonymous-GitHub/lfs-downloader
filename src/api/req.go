package api

import (
	"github.com/imroc/req/v3"
)

type HttpClient interface {
	Get(url string) (*req.Response, error)
	Post(url string, body interface{}) (*req.Response, error)
}
