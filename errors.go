package opnsense

import (
	"errors"
	"net/http"
)

const (
	ErrorEmptyCredentials = "invalid credentials: username & password must be specified"
	ErrorEmptyUsername    = "invalid credentials: username must be specified"
	ErrorEmptyPassword    = "invalid credentials: password must be specified"
	ErrorEmptyHost        = "invalid host: you must supply the host of the firewall"
	ErrorMissingOptions   = "missing api options: you must supply options for the api client"

	// http
	HttpError400 = "bad request: something went wrong. enable debug logging for more information"
	HttpError401 = "unauthorized: username or password is bad or you are not authorized to access this resource"
)

func httpStatusCheck(resp *http.Response) error {
	switch resp.StatusCode {
	case 400:
		return errors.New(HttpError400)
	case 401:
		return errors.New(HttpError401)
	}
	return nil
}
