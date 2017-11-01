package common

import (
	"encoding/hex"
	"encoding/json"
)

const (
	ConnectorUnknown     = 0x00
	ConnectorSc          = 0x01
	ConnectorFcStyle1    = 0x02
	ConnectorFcStyle2    = 0x03
	ConnectorBncTnc      = 0x04
	ConnectorFcCoax      = 0x05
	ConnectorFiberJack   = 0x06
	ConnectorLc          = 0x07
	ConnectorMtRj        = 0x08
	ConnectorMu          = 0x09
	ConnectorSg          = 0x0A
	ConnectorOptPtail    = 0x0B
	ConnectorMpo         = 0x0C
	ConnectorMpo2        = 0x0D
	ConnectorHssdcII     = 0x20
	ConnectorCopperPtail = 0x21
	ConnectorRj45        = 0x22
	ConnectorNoSeparable = 0x23
	ConnectorMxc2x16     = 0x24
)

var connectorNames = map[byte]string{
	ConnectorUnknown:     "Unknown or unspecified",
	ConnectorSc:          "SC",
	ConnectorFcStyle1:    "Fibre Channel style 1 copper",
	ConnectorFcStyle2:    "Fibre Channel style 2 copper",
	ConnectorBncTnc:      "BNC/TNC",
	ConnectorFcCoax:      "Fibre Channel coaxial headers",
	ConnectorFiberJack:   "FibreJack",
	ConnectorLc:          "LC",
	ConnectorMtRj:        "MT-RJ",
	ConnectorMu:          "MU",
	ConnectorSg:          "SG",
	ConnectorOptPtail:    "Optical pigtail",
	ConnectorMpo:         "MPO Parallel Optic",
	ConnectorMpo2:        "MPO Parallel Optic - 2x16",
	ConnectorHssdcII:     "HSSDC II",
	ConnectorCopperPtail: "Copper pigtail",
	ConnectorRj45:        "RJ45",
	ConnectorNoSeparable: "No separable connector",
	ConnectorMxc2x16:     "MXC 2x16",
}

type Connector byte

func (c Connector) String() string {
	n, ok := connectorNames[byte(c)]
	if !ok {
		return "Reserved or unknown"
	}
	return n
}

func (c Connector) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": c.String(),
		"hex":   hex.EncodeToString([]byte{byte(c)}),
	}
	return json.Marshal(m)
}

func (c *Connector) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*c = Connector(b[0])
	return nil
}
