package opnsense

import (
	"encoding/json"
	"regexp"
)

const (
	diagnosticsInterfaceBasePath               = "diagnostics/interface/"
	diagnosticsInterfaceDeleteRoute            = diagnosticsInterfaceBasePath + "delRoute"
	diagnosticsInterfaceFlushArp               = diagnosticsInterfaceBasePath + "flushArp"
	diagnosticsInterfaceGetArp                 = diagnosticsInterfaceBasePath + "getArp"
	diagnosticsInterfaceGetBpfStatistics       = diagnosticsInterfaceBasePath + "getBpfStatistics"
	diagnosticsInterfaceGetInterfaceNames      = diagnosticsInterfaceBasePath + "getInterfaceNames"
	diagnosticsInterfaceGetInterfaceStatistics = diagnosticsInterfaceBasePath + "getInterfaceStatistics"
	diagnosticsInterfaceGetMemoryStatistics    = diagnosticsInterfaceBasePath + "getMemoryStatistics"
	diagnosticsInterfaceGetNdp                 = diagnosticsInterfaceBasePath + "getNdp"
	diagnosticsInterfaceGetGetNetisrStatistics = diagnosticsInterfaceBasePath + "getNetisrStatistics"
	diagnosticsInterfaceGetProtocolStatistics  = diagnosticsInterfaceBasePath + "getProtocolStatistics"
	diagnosticsInterfaceGetRoutes              = diagnosticsInterfaceBasePath + "getRoutes"
	diagnosticsInterfaceGetSocketStatistics    = diagnosticsInterfaceBasePath + "getSocketStatistics"
)

type DiagnosticsInterfaceArp struct {
	Mac             string `json:"mac"`
	IP              string `json:"ip"`
	Intf            string `json:"intf"`
	Expired         bool   `json:"expired"`
	Expires         int    `json:"expires"`
	Permanent       bool   `json:"permanent"`
	Type            string `json:"type"`
	Manufacturer    string `json:"manufacturer"`
	Hostname        string `json:"hostname"`
	IntfDescription string `json:"intf_description"`
}

type DiagnosticsInterfaceBpfStatistics struct {
	BpfStatistics struct {
		BpfEntry []struct {
			Pid               int    `json:"pid"`
			InterfaceName     string `json:"interface-name"`
			Immediate         bool   `json:"immediate,omitempty"`
			Direction         string `json:"direction"`
			ReceivedPackets   int    `json:"received-packets"`
			DroppedPackets    int    `json:"dropped-packets"`
			FilterPackets     int    `json:"filter-packets"`
			StoreBufferLength int    `json:"store-buffer-length"`
			HoldBufferLength  int    `json:"hold-buffer-length"`
			Process           string `json:"process"`
			Promiscuous       bool   `json:"promiscuous,omitempty"`
			HeaderComplete    bool   `json:"header-complete,omitempty"`
			Locked            bool   `json:"locked,omitempty"`
		} `json:"bpf-entry"`
	} `json:"bpf-statistics"`
}

type DiagnosticsInterface struct {
	Name            string
	Interface       string `json:"name"`
	Flags           string `json:"flags"`
	Mtu             int    `json:"mtu"`
	Network         string `json:"network"`
	Address         string `json:"address"`
	ReceivedPackets int    `json:"received-packets"`
	ReceivedErrors  int    `json:"received-errors"`
	DroppedPackets  int    `json:"dropped-packets"`
	ReceivedBytes   int64  `json:"received-bytes"`
	SentPackets     int    `json:"sent-packets"`
	SendErrors      int    `json:"send-errors"`
	SentBytes       int64  `json:"sent-bytes"`
	Collisions      int    `json:"collisions"`
}

type DiagnosticsInterfaceStatistics map[string]map[string]DiagnosticsInterface

func (api *Api) DiagnosticsInterfaceArp() (*[]DiagnosticsInterfaceArp, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetArp, nil)
	if err != nil {
		return nil, err
	}

	arp := []DiagnosticsInterfaceArp{}

	err = json.Unmarshal(b, &arp)
	if err != nil {
		return nil, err
	}

	return &arp, nil
}

func (api *Api) DiagnosticsInterfaceBpfStatistics() (*DiagnosticsInterfaceBpfStatistics, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetBpfStatistics, nil)
	if err != nil {
		return nil, err
	}

	bpfStatistics := DiagnosticsInterfaceBpfStatistics{}

	err = json.Unmarshal(b, &bpfStatistics)
	if err != nil {
		return nil, err
	}

	return &bpfStatistics, nil
}

func (api *Api) DiagnosticsInterfaceStatistics() (*[]DiagnosticsInterface, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetInterfaceStatistics, nil)
	if err != nil {
		return nil, err
	}

	interfaceStatistics := DiagnosticsInterfaceStatistics{}

	err = json.Unmarshal(b, &interfaceStatistics)
	if err != nil {
		return nil, err
	}

	s := interfaceStatistics["statistics"]
	var interfaces []DiagnosticsInterface

	// why do you put data in a key? come on....
	// the keys come back like this "[WAN] (igb0) 00:00:00:00:00:00"... why do that? dynamic keys are garbage
	// which requires garbage
	r := regexp.MustCompile(`\[([^\[\]]*)\]`)
	for i := range s {
		// I need to make a copy of the original interface so that I can update its name... because again,
		// dynamic keys can shove it
		var t DiagnosticsInterface = s[i]
		name := r.FindAllStringSubmatch(i, -1)
		t.Name = name[0][1]

		interfaces = append(interfaces, t)
	}

	return &interfaces, nil
}
