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

type DiagnosticsInterfaceBpfStatisticsEntry struct {
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
}

type DiagnosticsInterfaceBpfStatistics map[string]map[string][]DiagnosticsInterfaceBpfStatisticsEntry

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

type DiagnosticsInterfaceMemoryStatisticsEntry struct {
	MbufCurrent                int `json:"mbuf-current"`
	MbufCache                  int `json:"mbuf-cache"`
	MbufTotal                  int `json:"mbuf-total"`
	ClusterCurrent             int `json:"cluster-current"`
	ClusterCache               int `json:"cluster-cache"`
	ClusterTotal               int `json:"cluster-total"`
	ClusterMax                 int `json:"cluster-max"`
	PacketCount                int `json:"packet-count"`
	PacketFree                 int `json:"packet-free"`
	JumboCount                 int `json:"jumbo-count"`
	JumboCache                 int `json:"jumbo-cache"`
	JumboTotal                 int `json:"jumbo-total"`
	JumboMax                   int `json:"jumbo-max"`
	JumboPageSize              int `json:"jumbo-page-size"`
	Jumbo9Count                int `json:"jumbo9-count"`
	Jumbo9Cache                int `json:"jumbo9-cache"`
	Jumbo9Total                int `json:"jumbo9-total"`
	Jumbo9Max                  int `json:"jumbo9-max"`
	Jumbo16Count               int `json:"jumbo16-count"`
	Jumbo16Cache               int `json:"jumbo16-cache"`
	Jumbo16Total               int `json:"jumbo16-total"`
	Jumbo16Limit               int `json:"jumbo16-limit"`
	BytesInUse                 int `json:"bytes-in-use"`
	BytesInCache               int `json:"bytes-in-cache"`
	BytesTotal                 int `json:"bytes-total"`
	MbufFailures               int `json:"mbuf-failures"`
	ClusterFailures            int `json:"cluster-failures"`
	PacketFailures             int `json:"packet-failures"`
	MbufSleeps                 int `json:"mbuf-sleeps"`
	ClusterSleeps              int `json:"cluster-sleeps"`
	PacketSleeps               int `json:"packet-sleeps"`
	JumbopSleeps               int `json:"jumbop-sleeps"`
	Jumbo9Sleeps               int `json:"jumbo9-sleeps"`
	Jumbo16Sleeps              int `json:"jumbo16-sleeps"`
	JumbopFailures             int `json:"jumbop-failures"`
	Jumbo9Failures             int `json:"jumbo9-failures"`
	Jumbo16Failures            int `json:"jumbo16-failures"`
	SendfileSyscalls           int `json:"sendfile-syscalls"`
	SendfileNoIo               int `json:"sendfile-no-io"`
	SendfileIoCount            int `json:"sendfile-io-count"`
	SendfilePagesSent          int `json:"sendfile-pages-sent"`
	SendfilePagesValid         int `json:"sendfile-pages-valid"`
	SendfilePagesBogus         int `json:"sendfile-pages-bogus"`
	SendfileRequestedReadahead int `json:"sendfile-requested-readahead"`
	SendfileReadahead          int `json:"sendfile-readahead"`
	SendfileBusyEncounters     int `json:"sendfile-busy-encounters"`
	SfbufsAllocFailed          int `json:"sfbufs-alloc-failed"`
	SfbufsAllocWait            int `json:"sfbufs-alloc-wait"`
}

type DiagnosticsInterfaceMemoryStatistics map[string]DiagnosticsInterfaceMemoryStatisticsEntry

type DiagnosticsInterfaceNdp struct {
	Mac             string `json:"mac"`
	IP              string `json:"ip"`
	Intf            string `json:"intf"`
	Manufacturer    string `json:"manufacturer"`
	IntfDescription string `json:"intf_description"`
}

type DiagnosticsInterfaceRoute struct {
	Proto           string `json:"proto"`
	Destination     string `json:"destination"`
	Gateway         string `json:"gateway"`
	Flags           string `json:"flags"`
	Use             string `json:"use"`
	Mtu             string `json:"mtu"`
	Netif           string `json:"netif"`
	Expire          string `json:"expire"`
	IntfDescription string `json:"intf_description"`
}

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

func (api *Api) DiagnosticsInterfaceRoutes() (*[]DiagnosticsInterfaceRoute, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetRoutes, nil)
	if err != nil {
		return nil, err
	}

	routes := []DiagnosticsInterfaceRoute{}

	err = json.Unmarshal(b, &routes)
	if err != nil {
		return nil, err
	}

	return &routes, nil
}

func (api *Api) DiagnosticsInterfaceBpfStatistics() (*[]DiagnosticsInterfaceBpfStatisticsEntry, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetBpfStatistics, nil)
	if err != nil {
		return nil, err
	}

	bpfStatistics := DiagnosticsInterfaceBpfStatistics{}

	err = json.Unmarshal(b, &bpfStatistics)
	if err != nil {
		return nil, err
	}

	s := bpfStatistics["bpf-statistics"]["bpf-entry"]

	return &s, nil
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

func (api *Api) DiagnosticsInterfaceMemoryStatistics() (*DiagnosticsInterfaceMemoryStatisticsEntry, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetMemoryStatistics, nil)
	if err != nil {
		return nil, err
	}

	interfaceMemoryStatistics := DiagnosticsInterfaceMemoryStatistics{}

	err = json.Unmarshal(b, &interfaceMemoryStatistics)
	if err != nil {
		return nil, err
	}

	s := interfaceMemoryStatistics["mbuf-statistics"]

	return &s, nil
}

func (api *Api) DiagnosticsInterfaceNdp() (*[]DiagnosticsInterfaceNdp, error) {
	b, err := api.Do("GET", api.Host+diagnosticsInterfaceGetNdp, nil)
	if err != nil {
		return nil, err
	}

	interfaceNdp := []DiagnosticsInterfaceNdp{}

	err = json.Unmarshal(b, &interfaceNdp)
	if err != nil {
		return nil, err
	}

	return &interfaceNdp, nil
}
