package entities

type Apn struct {
	Text               string `xml:",chardata"`
	Carrier            string `xml:"carrier,attr"`
	CarrierID          string `xml:"carrier_id,attr"`
	Mcc                string `xml:"mcc,attr"`
	Mnc                string `xml:"mnc,attr"`
	Apn                string `xml:"apn,attr"`
	User               string `xml:"user,attr"`
	Password           string `xml:"password,attr"`
	Proxy              string `xml:"proxy,attr"`
	Port               string `xml:"port,attr"`
	Mmsc               string `xml:"mmsc,attr"`
	Mmsproxy           string `xml:"mmsproxy,attr"`
	Mmsport            string `xml:"mmsport,attr"`
	Type               string `xml:"type,attr"`
	Mtu                string `xml:"mtu,attr"`
	Protocol           string `xml:"protocol,attr"`
	RoamingProtocol    string `xml:"roaming_protocol,attr"`
	BearerBitmask      string `xml:"bearer_bitmask,attr"`
	ProfileID          string `xml:"profile_id,attr"`
	ModemCognitive     string `xml:"modem_cognitive,attr"`
	MaxConns           string `xml:"max_conns,attr"`
	MaxConnsTime       string `xml:"max_conns_time,attr"`
	CarrierEnabled     string `xml:"carrier_enabled,attr"`
	Authtype           string `xml:"authtype,attr"`
	MvnoMatchData      string `xml:"mvno_match_data,attr"`
	MvnoType           string `xml:"mvno_type,attr"`
	Server             string `xml:"server,attr"`
	UserVisible        string `xml:"user_visible,attr"`
	Bearer             string `xml:"bearer,attr"`
	Mtusize            string `xml:"mtusize,attr"`
	NetworkTypeBitmask string `xml:"network_type_bitmask,attr"`
}