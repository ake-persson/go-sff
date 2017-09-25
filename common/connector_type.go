package common

import (
	"encoding/hex"
	"encoding/json"
)

const (
	ConnectorTypeUnknown     = 0x00
	ConnectorTypeSC          = 0x01
	ConnectorTypeFCStyle1    = 0x02
	ConnectorTypeFCStyle2    = 0x03
	ConnectorTypeBNCTNC      = 0x04
	ConnectorTypeFCCoax      = 0x05
	ConnectorTypeFiberJack   = 0x06
	ConnectorTypeLC          = 0x07
	ConnectorTypeMTRJ        = 0x08
	ConnectorTypeMU          = 0x09
	ConnectorTypeSG          = 0x0A
	ConnectorTypeOptPt       = 0x0B
	ConnectorTypeMPO         = 0x0C
	ConnectorTypeMPO2        = 0x0D
	ConnectorTypeHSSDCII     = 0x20
	ConnectorTypeCopperPt    = 0x21
	ConnectorTypeRJ45        = 0x22
	ConnectorTypeNoSeparable = 0x23
	ConnectorTypeMXC2x16     = 0x24
)

var connectorNames = map[byte]string{
	ConnectorTypeUnknown:     "Unknown or unspecified",
	ConnectorTypeSC:          "SC",
	ConnectorTypeFCStyle1:    "Fibre Channel style 1 copper",
	ConnectorTypeFCStyle2:    "Fibre Channel style 2 copper",
	ConnectorTypeBNCTNC:      "BNC/TNC",
	ConnectorTypeFCCoax:      "Fibre Channel coaxial headers",
	ConnectorTypeFiberJack:   "FibreJack",
	ConnectorTypeLC:          "LC",
	ConnectorTypeMTRJ:        "MT-RJ",
	ConnectorTypeMU:          "MU",
	ConnectorTypeSG:          "SG",
	ConnectorTypeOptPt:       "Optical pigtail",
	ConnectorTypeMPO:         "MPO Parallel Optic",
	ConnectorTypeMPO2:        "MPO Parallel Optic - 2x16",
	ConnectorTypeHSSDCII:     "HSSDC II",
	ConnectorTypeCopperPt:    "Copper pigtail",
	ConnectorTypeRJ45:        "RJ45",
	ConnectorTypeNoSeparable: "No separable connector",
	ConnectorTypeMXC2x16:     "MXC 2x16",
}

type ConnectorType byte

func (c ConnectorType) String() string {
	n, ok := connectorNames[byte(c)]
	if !ok {
		return "Reserved or unknown"
	}
	return n
}

func (c ConnectorType) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": c.String(),
		"hex":  hex.EncodeToString([]byte{byte(c)}),
	}
	return json.Marshal(m)
}
