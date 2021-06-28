package opnsense

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 60
	defaultLogging = "info"
	apiPath        = "/api/"
)

// The API client
type Api struct {
	Username   string
	Password   string
	Host       string
	Options    *ApiOptions
	httpClient *http.Client
}

// API client options
type ApiOptions struct {
	IgnoreSslErrors bool
	TimeOut         int
	Logging         string
	Proxy           string
	Print           string
}

// Constructor for building API options with default values
func NewApiDefaultApiOptions() *ApiOptions {
	var o ApiOptions
	o.defaultOptions()
	return &o
}

// Build default API client options
func (o *ApiOptions) defaultOptions() {
	if o.TimeOut == 0 {
		o.TimeOut = defaultTimeout
	}
	if o.Logging == "" {
		o.Logging = defaultLogging
	}
}

// Build HTTP client options
func (api *Api) httpOptions() error {
	tr := &http.Transport{}

	if api.Options.IgnoreSslErrors {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if api.Options.Proxy != "" {
		if u, err := url.Parse(api.Options.Proxy); err != nil {
			return err
		} else {
			tr.Proxy = http.ProxyURL(u)
		}
	}

	api.httpClient.Transport = tr

	api.httpClient.Timeout = time.Duration(api.Options.TimeOut) * time.Second

	return nil
}

// Creates a API client that uses basic auth
func NewApiBasicAuth(username string, password string, host string, options *ApiOptions) (*Api, error) {
	if username == "" || password == "" {
		return nil, errors.New(ErrorEmptyCredentials)
	}

	if host == "" {
		return nil, errors.New(ErrorEmptyHost)
	}

	if options == nil {
		options = NewApiDefaultApiOptions()
	}

	api := &Api{
		Host:       "https://" + host + apiPath,
		Username:   username,
		Password:   password,
		Options:    options,
		httpClient: &http.Client{},
	}

	err := api.httpOptions()
	if err != nil {
		return nil, err
	}

	return api, nil
}

// A general Do function for a API request
func (api *Api) Do(method, url string, body []byte) ([]byte, error) {
	r := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(api.Username, api.Password)
	req.Header.Add("ContentType", "application/json;charset=utf-8")

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := httpStatusCheck(resp); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return b, nil
}
