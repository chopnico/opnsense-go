package opnsense

import (
	"encoding/json"
	"errors"
	"net/url"
)

type Route struct {
	UUID     string `json:"uuid"`
	Disabled string `json:"disabled"`
	Network  string `json:"network"`
	Gateway  string `json:"gateway"`
	Descr    string `json:"descr"`
}

type savedRoute struct {
	result string `json:"result"`
	uuid   string `json:"uuid"`
}
type routes struct {
	Rows []Route `json:"rows"`
}

// retrieve all routes
func (api *Api) GetRoutes() (*[]Route, error) {
	b, err := api.Do("GET", "/routes/routes/searchroute", nil)
	if err != nil {
		return nil, err
	}

	routes := routes{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	return &routes.Rows, nil
}

// retrieve a single route using its uuid
func (api *Api) GetRouteByUuid(uuid string) (*Route, error) {
	uuid = url.PathEscape(uuid)
	b, err := api.Do("GET", "/routes/routes/searchroute", nil)
	if err != nil {
		return nil, err
	}

	routes := routes{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	// we're assuming that the uuid is unique
	var route Route
	for _, i := range routes.Rows {
		if i.UUID == uuid {
			route = i
		}
	}

	return &route, nil
}

// create a route
func (api *Api) SetRoute(r Route) (*Route, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	b, err := api.Do("POST", "/routes/routes/addroute", body)
	if err != nil {
		return nil, err
	}

	resp := savedRoute{}
	if resp.result != "saved" {
		return nil, errors.New("failed to add route, having status of : " + resp.result + ". enable debugging for more information.")
	}

	routes := routes{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	// we're assuming that the uuid is unique
	var route Route
	for _, i := range routes.Rows {
		if i.UUID == resp.uuid {
			route = i
		}
	}

	return &route, nil
}
