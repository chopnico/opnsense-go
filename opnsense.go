package opnsense

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/chopnico/output"
)

const (
	defaultTimeout = 60
	defaultLogging = "info"
	apiPath        = "/api"
)

// The API client
type Api struct {
	options    map[string]interface{}
	httpClient *http.Client
}

type apiResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// REQUIRED: sets username
func (api Api) username(v string) Api {
	api.option("username", v)
	return api
}

// REQUIRED: sets password
func (api Api) password(v string) Api {
	api.option("password", v)
	return api
}

// REQUIRED: sets host of firewall endpoint
func (api Api) url(v string) Api {
	api.option("url", v)
	return api
}

// sets timeout
func (api Api) Timeout(v int) Api {
	api.option("timeout", v)
	return api
}

// sets whether the http client should ignore ssl errors
func (api Api) IgnoreSslErrors() Api {
	api.options["ignore-ssl"] = true
	api.httpOptions()
	return api
}

// sets the proxy for the http client
func (api Api) Proxy(v string) Api {
	api.options["proxy"] = v
	api.httpOptions()
	return api
}

// sets the info logger
func (api Api) InfoLogger(v *log.Logger) Api {
	api.option("logger-info", v)
	return api
}

// sets the debug logger
func (api Api) DebugLogger(v *log.Logger) Api {
	api.option("logger-debug", v)
	return api
}

// sets the logging level
func (api Api) LoggingLevel(v string) Api {
	api.option("logging-level", v)
	return api
}

func (api *Api) option(k string, v interface{}) {
	if api.options == nil {
		api.options = make(map[string]interface{})
		api.option("logging-level", defaultLogging)
		api.option("timeout", defaultTimeout)
		api.option("ignore-ssl", false)
		api.option("proxy", "")
	}
	api.options[k] = v
}

// Build HTTP client options
func (api Api) httpOptions() {
	tr := &http.Transport{}

	if api.options["ignore-ssl"].(bool) {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if api.options["proxy"].(string) != "" {
		u, _ := url.Parse(api.options["proxy"].(string))
		tr.Proxy = http.ProxyURL(u)
	}

	api.httpClient.Transport = tr

	api.httpClient.Timeout = time.Duration(api.options["timeout"].(int)) * time.Second
}

// create the default debug logger
func defaultInfoLogger() *log.Logger {
	return log.New(os.Stderr, "[INFO] ", log.LstdFlags)
}

// create the default debug logger
func defaultDebugLogger() *log.Logger {
	return log.New(os.Stderr, "[DEBUG] ", log.LstdFlags)
}

// creates a response error
func newResponseError(b []byte) apiResponseError {
	r := apiResponseError{}
	json.Unmarshal(b, &r)
	return r
}

func (api *Api) isLoggingDebug() bool {
	if api.options["logging-level"].(string) == "debug" {
		return true
	}
	return false
}

func (api *Api) isLoggingInfo() bool {
	if api.options["logging-level"].(string) == "info" {
		return true
	}
	return false
}

// Creates a API client that uses basic auth
func NewApiBasicAuth(username string, password string, host string) (Api, error) {
	api := Api{
		httpClient: &http.Client{},
	}

	api.option("username", username)
	api.option("password", password)
	api.option("url", "https://"+host+apiPath)
	api.option("logger-info", defaultInfoLogger())
	api.option("logger-debug", defaultDebugLogger())

	return api, nil
}

// The main do function for an api request
func (api *Api) Do(method, path string, body []byte) ([]byte, error) {
	r := bytes.NewReader(body)

	req, err := http.NewRequest(method, api.options["url"].(string)+path, r)
	if api.isLoggingDebug() {
		api.WriteToDebugLog("request url : " + req.URL.Host)
		api.WriteToDebugLog("request method : " + req.Method)
		api.WriteToDebugLog("request headers : " + output.FormatItemAsPrettyJson(req.Header))
	}
	if err != nil {
		if api.isLoggingDebug() {
			return nil, errors.New("debugging")
		}
		return nil, err
	}

	req.SetBasicAuth(api.options["username"].(string), api.options["password"].(string))
	req.Header.Add("ContentType", "application/json;charset=utf-8")

	resp, err := api.httpClient.Do(req)
	if api.isLoggingDebug() {
		api.WriteToDebugLog("request headers : " + output.FormatItemAsPrettyJson(req.Header))
		api.WriteToDebugLog("request uri : " + req.URL.RequestURI())
		api.WriteToDebugLog("request method : " + req.Method)
	}
	if err != nil {
		if api.isLoggingDebug() {
			return nil, errors.New("debugging")
		}
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if api.isLoggingDebug() {
			return nil, errors.New("debugging")
		}
		return nil, err
	}

	defer resp.Body.Close()

	if api.isLoggingDebug() {
		if err != nil {
			e := newResponseError(b)
			api.WriteToDebugLog("response message : " + e.Message)
			api.WriteToDebugLog("response status code : " + fmt.Sprintf("%d", e.Status))
		}
		api.WriteToDebugLog("request headers : " + output.FormatItemAsPrettyJson(req.Header))
		api.WriteToDebugLog("request uri : " + req.URL.RequestURI())
		api.WriteToDebugLog("request method : " + req.Method)
		api.WriteToDebugLog("response headers : " + output.FormatItemAsPrettyJson(resp.Header))
	}

	if resp.StatusCode != 200 {
		if api.isLoggingDebug() {
			return nil, errors.New("debugging")
		}
		return nil, errors.New("something went wrong. enable debug logs for more information")
	}

	return b, nil
}

func (api *Api) WriteToDebugLog(msg string) {
	logger := api.options["logger-debug"].(*log.Logger)
	logger.Println(msg)
}

func (api *Api) WriteToInfoLog(msg string) {
	logger := api.options["logger-info"].(*log.Logger)
	logger.Println(msg)
}
