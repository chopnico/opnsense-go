package opnsense

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/google/uuid"
)

// a route
type Route struct {
	UUID        string `json:"uuid,omitempty"`
	Disabled    string `json:"disabled"`
	Network     string `json:"network"`
	Gateway     string `json:"gateway"`
	Description string `json:"descr"`
}

// the response returned for deleting, saving,
// and creating routes
type responseRoute struct {
	Uuid        string `json:"uuid,omitempty"`
	Result      string `json:"result"`
	Validations struct {
		RouteNetwork string `json:"route.network"`
	} `json:"validations,omitempty"`
}

// the api requires a route to be wrapped
// before sending
type newRoute struct {
	Route Route `json:"route"`
}

// the api returns the results wrapped in
// a "rows" object
type Routes struct {
	List []Route `json:"rows"`
}

// retrieve all routes
func (api *Api) GetRoutes(enabled bool) (*[]Route, error) {
	// set http headers for request
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json; charset=UTF-8"
	api.HttpHeaders(headers)

	// perform the api request
	b, err := api.Do("GET", "/routes/routes/searchroute", nil)
	if err != nil {
		return nil, err
	}

	// create an empty list of routes
	// unmarshal the bytes
	// return the routes
	routes := Routes{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	if enabled {
		r := []Route{}
		for _, i := range routes.List {
			if i.Disabled == "0" {
				r = append(r, i)
			}
		}
		return &r, nil
	} else {
		return &routes.List, nil
	}
}

// retrieve routes by network
func (api *Api) GetRoutesByNetwork(n string, enabled bool) (*[]Route, error) {
	routes, err := api.GetRoutes(enabled)
	if err != nil {
		return nil, err
	}

	r := []Route{}

	for _, i := range *routes {
		if i.Network == n {
			r = append(r, i)
		}
	}

	return &r, nil
}

// retrieve routes by network
func (api *Api) GetRoutesByNetworkWithDescription(n, d string, enabled bool) (*[]Route, error) {
	routes, err := api.GetRoutes(enabled)
	if err != nil {
		return nil, err
	}

	r := []Route{}

	for _, i := range *routes {
		if i.Network == n && i.Description == d {
			r = append(r, i)
		}
	}

	return &r, nil
}

// retrieve routes by destination
func (api *Api) GetRoutesByDescription(d string, enabled bool) (*[]Route, error) {
	routes, err := api.GetRoutes(enabled)
	if err != nil {
		return nil, err
	}

	r := []Route{}
	for _, i := range *routes {
		if i.Description == d {
			r = append(r, i)
		}
	}

	return &r, nil
}

// retrieve a single route using its uuid
func (api *Api) GetRouteByUuid(u string) (*Route, error) {
	_, err := uuid.Parse(u)
	if err != nil {
		return nil, errors.New("route id invalid")
	}

	// set http headers for request
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json; charset=UTF-8"
	api.HttpHeaders(headers)

	// make sure the path is clean
	u = url.PathEscape(u)
	// perform the api request
	b, err := api.Do("GET", "/routes/routes/searchroute", nil)
	if err != nil {
		return nil, err
	}

	// create an empty list of routes
	// unmarshal the bytes into the route list
	routes := Routes{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	// pull the route from the route list using the uuid given
	// we're assuming that the uuid is unique
	route := Route{}
	for _, i := range routes.List {
		if i.UUID == u {
			route = i
		}
	}

	return &route, nil
}

// delete a route using the uuid
func (api *Api) DeleteRoute(u string) error {
	_, err := uuid.Parse(u)
	if err != nil {
		return errors.New("route id invalid")
	}

	// set http headers for request
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json; charset=UTF-8"
	api.HttpHeaders(headers)

	// clean up the uuid before submitting
	u = url.PathEscape(u)
	b, err := api.Do("POST", "/routes/routes/delroute/"+u, nil)
	if err != nil {
		return err
	}

	// create the response object
	// unmarshal them bytes
	resp := responseRoute{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return err
	}

	// check if the result was deleted or not
	switch resp.Result {
	case "deleted":
		return nil
	case "not found":
		return errors.New("route not found")
	default:
		if resp.Validations.RouteNetwork != "" {
			return errors.New(resp.Validations.RouteNetwork)
		}
		return errors.New("failed to delete route, having status of : " + resp.Result + ". enable debugging for more information")
	}
}

// create a route
func (api *Api) SetRoute(r Route) (*Route, error) {
	// set http headers for request
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json, text/javascript, */*; q=0.01"
	api.HttpHeaders(headers)

	// create a new route, wrapped around a route hash
	// because... i don't know why
	nroute := newRoute{Route: r}

	// create the post body
	body, err := json.Marshal(nroute)
	if err != nil {
		return nil, err
	}

	// perform api request, saving them bytes
	b, err := api.Do("POST", "/routes/routes/addroute", body)
	if err != nil {
		return nil, err
	}

	// create a response route to check if it failed
	resp := responseRoute{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	switch resp.Result {
	case "saved":
		// retrieve the newly created route
		u, err := api.GetRouteByUuid(resp.Uuid)
		if err != nil {
			return nil, err
		}

		// unmarshal that route
		err = json.Unmarshal(b, &u)
		if err != nil {
			return nil, err
		}
		return u, nil
	default:
		if resp.Validations.RouteNetwork != "" {
			return nil, errors.New(resp.Validations.RouteNetwork)
		}
		return nil, errors.New("failed to add route, having status of : " + resp.Result + ". enable debugging for more information")
	}
}
