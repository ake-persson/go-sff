package common

import (
	"encoding/hex"
	"encoding/json"
)

const (
	IdentifierUnknown    = 0x00
	IdentifierGBIC       = 0x01
	IdentifierSoldered   = 0x02
	IdentifierSFP        = 0x03
	Identifier300PinXBI  = 0x04
	IdentifierXENPAK     = 0x05
	IdentifierXFP        = 0x06
	IdentifierXFF        = 0x07
	IdentifierXFPE       = 0x08
	IdentifierXPAK       = 0x09
	IdentifierX2         = 0x0A
	IdentifierDWDMSFP    = 0x0B
	IdentifierQSFP       = 0x0C
	IdentifierQSFPPlus   = 0x0D
	IdentifierCXP        = 0x0E
	IdentifierHD4X       = 0x0F
	IdentifierHD8X       = 0x10
	IdentifierQSFP28     = 0x11
	IdentifierCXP2       = 0x12
	IdentifierCDFP       = 0x13
	IdentifierHD4XFanout = 0x14
	IdentifierHD8XFanout = 0x15
	IdentifierCDFPStyle3 = 0x16
	IdentifierMicroQSFP  = 0x17
)

var identifierNames = map[byte]string{
	IdentifierUnknown:    "No module present, unknown, or unspecified",
	IdentifierGBIC:       "GBIC",
	IdentifierSoldered:   "Module soldered to motherboard",
	IdentifierSFP:        "SFP",
	Identifier300PinXBI:  "300 pin XBI",
	IdentifierXENPAK:     "XENPAK",
	IdentifierXFP:        "XFP",
	IdentifierXFF:        "XFF",
	IdentifierXFPE:       "XFP-E",
	IdentifierXPAK:       "XPAK",
	IdentifierX2:         "X2",
	IdentifierDWDMSFP:    "DWDM-SFP",
	IdentifierQSFP:       "QSFP",
	IdentifierQSFPPlus:   "QSFP+",
	IdentifierCXP:        "CXP",
	IdentifierHD4X:       "Shielded Mini Multilane HD 4X",
	IdentifierHD8X:       "Shielded Mini Multilane HD 8X",
	IdentifierQSFP28:     "QSFP28",
	IdentifierCXP2:       "CXP2/CXP28",
	IdentifierCDFP:       "CDFP Style 1/Style 2",
	IdentifierHD4XFanout: "Shielded Mini Multilane HD 4X Fanout Cable",
	IdentifierHD8XFanout: "Shielded Mini Multilane HD 8X Fanout Cable",
	IdentifierCDFPStyle3: "CDFP Style 3",
	IdentifierMicroQSFP:  "MicroQSFP",
}

type Identifier byte

func (i Identifier) String() string {
	s, ok := identifierNames[byte(i)]
	if !ok {
		return "Reserved or unknown"
	}
	return s
}

func (i Identifier) MarshalJSON() ([]byte, error) {
	b := []byte{byte(i)}
	m := map[string]interface{}{
		"name": i.String(),
		"hex":  hex.EncodeToString(b),
	}
	return json.Marshal(m)
}
