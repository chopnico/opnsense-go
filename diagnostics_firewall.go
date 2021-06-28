package opnsense

import (
	"encoding/json"
)

const (
	diagnosticsFirewallBasePath = "diagnostics/firewall/"
	diagnosticsFirewallLog      = diagnosticsFirewallBasePath + "log"
)

type DiagnosticsFirewallLogEntry struct {
	Rulenr        string   `json:"rulenr"`
	Subrulenr     string   `json:"subrulenr"`
	Anchorname    string   `json:"anchorname"`
	Ridentifier   string   `json:"ridentifier"`
	Interface     string   `json:"interface"`
	Reason        string   `json:"reason"`
	Action        string   `json:"action"`
	Dir           string   `json:"dir"`
	Version       string   `json:"version"`
	Tos           string   `json:"tos,omitempty"`
	Ecn           string   `json:"ecn,omitempty"`
	TTL           string   `json:"ttl,omitempty"`
	ID            string   `json:"id,omitempty"`
	Offset        string   `json:"offset,omitempty"`
	Ipflags       string   `json:"ipflags,omitempty"`
	Proto         string   `json:"proto"`
	Protoname     string   `json:"protoname"`
	Length        string   `json:"length,omitempty"`
	Src           string   `json:"src"`
	Dst           string   `json:"dst"`
	Srcport       string   `json:"srcport,omitempty"`
	Dstport       string   `json:"dstport,omitempty"`
	Datalen       string   `json:"datalen,omitempty"`
	Digest        string   `json:"__digest__"`
	Host          string   `json:"__host__"`
	Timestamp     string   `json:"__timestamp__"`
	Spec          []string `json:"__spec__"`
	Label         string   `json:"label"`
	Rid           string   `json:"rid"`
	Tcpflags      string   `json:"tcpflags,omitempty"`
	Seq           string   `json:"seq,omitempty"`
	Ack           string   `json:"ack,omitempty"`
	Urp           string   `json:"urp,omitempty"`
	Tcpopts       string   `json:"tcpopts,omitempty"`
	Class         string   `json:"class,omitempty"`
	Flowlabel     string   `json:"flowlabel,omitempty"`
	Hlim          string   `json:"hlim,omitempty"`
	PayloadLength string   `json:"payload-length,omitempty"`
}

func (api *Api) DiagnosticsFirewallLog(limit string) (*[]DiagnosticsFirewallLogEntry, error) {
	b, err := api.Do("GET", api.Host+diagnosticsFirewallLog+"/?limit="+limit, nil)
	if err != nil {
		return nil, err
	}

	logs := []DiagnosticsFirewallLogEntry{}

	err = json.Unmarshal(b, &logs)
	if err != nil {
		return nil, err
	}

	return &logs, nil
}
