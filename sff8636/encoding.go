package sff8636

import (
	"encoding/hex"
	"encoding/json"
)

const (
	EncodingUnspecified = 0x00
	Encoding8B10B       = 0x01
	Encoding4B5B        = 0x02
	EncodingNRZ         = 0x03
	Encoding4h          = 0x04
	Encoding5h          = 0x05
	Encoding6h          = 0x06
	Encoding256B        = 0x07
	EncodingPAM4        = 0x08
)

var encodingNames = map[byte]string{
	EncodingUnspecified: "Unspecified",
	Encoding8B10B:       "8B/10B",
	Encoding4B5B:        "4B/5B",
	EncodingNRZ:         "NRZ",
	Encoding4h:          "SONET Scrambled",
	Encoding5h:          "64B/66B",
	Encoding6h:          "Manchester",
	Encoding256B:        "256B/257B (transcoded FEC-enabled data)",
	EncodingPAM4:        "PAM4",
}

type Encoding byte

func (e Encoding) String() string {
	n, ok := encodingNames[byte(e)]
	if !ok {
		return "Reserved or unknown"
	}
	return n
}

func (e Encoding) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": e.String(),
		"hex":  hex.EncodeToString([]byte{byte(e)}),
	}
	return json.Marshal(m)
}
