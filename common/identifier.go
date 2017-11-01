package common

import (
	"encoding/hex"
	"encoding/json"
)

const (
	IdentifierUnknown    = 0x00
	IdentifierGbic       = 0x01
	IdentifierSoldered   = 0x02
	IdentifierSfp        = 0x03
	Identifier300PinXbi  = 0x04
	IdentifierXenpak     = 0x05
	IdentifierXfp        = 0x06
	IdentifierXff        = 0x07
	IdentifierXfpE       = 0x08
	IdentifierXpak       = 0x09
	IdentifierX2         = 0x0A
	IdentifierDwdmSfp    = 0x0B
	IdentifierQsfp       = 0x0C
	IdentifierQsfpPlus   = 0x0D
	IdentifierCxp        = 0x0E
	IdentifierHd4x       = 0x0F
	IdentifierHd8x       = 0x10
	IdentifierQsfp28     = 0x11
	IdentifierCxp2       = 0x12
	IdentifierCdfp       = 0x13
	IdentifierHd4xFanout = 0x14
	IdentifierHd8xFanout = 0x15
	IdentifierCdfpStyle3 = 0x16
	IdentifierMicroQsfp  = 0x17
)

var identifierNames = map[byte]string{
	IdentifierUnknown:    "No module present, unknown, or unspecified",
	IdentifierGbic:       "GBIC",
	IdentifierSoldered:   "Module soldered to motherboard",
	IdentifierSfp:        "SFP",
	Identifier300PinXbi:  "300 pin XBI",
	IdentifierXenpak:     "XENPAK",
	IdentifierXfp:        "XFP",
	IdentifierXff:        "XFF",
	IdentifierXfpE:       "XFP-E",
	IdentifierXpak:       "XPAK",
	IdentifierX2:         "X2",
	IdentifierDwdmSfp:    "DWDM-SFP",
	IdentifierQsfp:       "QSFP",
	IdentifierQsfpPlus:   "QSFP+",
	IdentifierCxp:        "CXP",
	IdentifierHd4x:       "Shielded Mini Multilane HD 4X",
	IdentifierHd8x:       "Shielded Mini Multilane HD 8X",
	IdentifierQsfp28:     "QSFP28",
	IdentifierCxp2:       "CXP2/CXP28",
	IdentifierCdfp:       "CDFP Style 1/Style 2",
	IdentifierHd4xFanout: "Shielded Mini Multilane HD 4X Fanout Cable",
	IdentifierHd8xFanout: "Shielded Mini Multilane HD 8X Fanout Cable",
	IdentifierCdfpStyle3: "CDFP Style 3",
	IdentifierMicroQsfp:  "MicroQSFP",
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
		"value": i.String(),
		"hex":   hex.EncodeToString(b),
	}
	return json.Marshal(m)
}

func (i *Identifier) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	v := Identifier(b[0])
	i = &v
	return nil
}
