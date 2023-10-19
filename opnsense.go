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

// the api client
type Api struct {
	options    map[string]interface{}
	httpClient *http.Client
}

// api error
type apiError struct {
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

// sets http client headers
func (api Api) HttpHeaders(v map[string]interface{}) Api {
	api.option("http-headers", v)
	return api
}

// debug writer
func (api *Api) WriteToDebugLog(msg string) {
	logger := api.options["logger-debug"].(*log.Logger)
	logger.Println(msg)
}

// info writer
func (api *Api) WriteToInfoLog(msg string) {
	logger := api.options["logger-info"].(*log.Logger)
	logger.Println(msg)
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
func newApiError(b []byte) apiError {
	r := apiError{}
	json.Unmarshal(b, &r)
	return r
}

// checks to see if logging level is set to debug
func (api *Api) isLoggingDebug() bool {
	if api.options["logging-level"].(string) == "debug" {
		return true
	}
	return false
}

// checks to see if logging level is set to info
func (api *Api) isLoggingInfo() bool {
	if api.options["logging-level"].(string) == "info" {
		return true
	}
	return false
}

func (api *Api) getHttpHeaders() map[string]interface{} {
	return api.options["http-headers"].(map[string]interface{})
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
	api.option("http-headers", make(map[string]interface{}))

	return api, nil
}

// The main do function for an api request
// lots of debugging code going on here
// REVIEW: refactor
func (api *Api) Do(method, path string, body []byte) ([]byte, error) {
	// read them bytes
	r := bytes.NewReader(body)

	// generate a new request
	req, err := http.NewRequest(method, api.options["url"].(string)+path, r)
	// DEBUG
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

	// set username and password for request
	req.SetBasicAuth(api.options["username"].(string), api.options["password"].(string))
	// set some headers for the request
	// headers are set using the api.HttpHeaders option
	headers := api.getHttpHeaders()
	if len(headers) != 0 {
		for k, v := range headers {
			req.Header.Set(k, fmt.Sprintf("%s", v))
		}
	} else {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	// make the formal request using the an http client
	resp, err := api.httpClient.Do(req)
	// DEBUG
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

	// read the binary response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if api.isLoggingDebug() {
			return nil, errors.New("debugging")
		}
		return nil, err
	}

	defer resp.Body.Close()

	//DEBUG
	if api.isLoggingDebug() {
		if err != nil {
			e := newApiError(b)
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
			e := newApiError(b)
			api.WriteToDebugLog("response message : " + e.Message)
			api.WriteToDebugLog("response status code : " + fmt.Sprintf("%d", e.Status))
			return nil, errors.New("debugging")
		}
		return nil, errors.New("something went wrong. enable debug logs for more information")
	}

	return b, nil
}
