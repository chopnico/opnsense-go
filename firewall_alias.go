package opnsense

import (
	"encoding/json"
)

const (
	firewallAliasBasePath = "firewall/alias/"
	firewallAliasExport   = firewallAliasBasePath + "export"
)

type FirewallAlias struct {
	Id          string
	Enabled     string `json:"enabled"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Proto       string `json:"proto"`
	Counters    string `json:"counters"`
	Updatefreq  string `json:"updatefreq"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type FirewallAliases struct {
	Aliases map[string]map[string]FirewallAlias
}

func (api *Api) FirewallAlias() (*[]FirewallAlias, error) {
	b, err := api.Do("GET", api.Host+firewallAliasExport, nil)
	if err != nil {
		return nil, err
	}

	aliasExport := FirewallAliases{}

	err = json.Unmarshal(b, &aliasExport)
	if err != nil {
		return nil, err
	}

	a := aliasExport.Aliases["alias"]
	var aliases []FirewallAlias

	for i := range a {
		var t FirewallAlias = a[i]
		t.Id = i
		aliases = append(aliases, t)
	}

	return &aliases, nil
}
