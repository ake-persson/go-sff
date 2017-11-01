package sff8079

import (
	"encoding/hex"
	"encoding/json"
)

const (
	EncodingUnspecified = 0x00
	Encoding8b10b       = 0x01
	Encoding4b5b        = 0x02
	EncodingNrz         = 0x03
	Encoding4h          = 0x04
	Encoding5h          = 0x05
	Encoding6h          = 0x06
	Encoding256b        = 0x07
	EncodingPam4        = 0x08
)

var encodingNames = map[byte]string{
	EncodingUnspecified: "Unspecified",
	Encoding8b10b:       "8B/10B",
	Encoding4b5b:        "4B/5B",
	EncodingNrz:         "NRZ",
	Encoding4h:          "Manchester",
	Encoding5h:          "SONET Scrambled",
	Encoding6h:          "64B/66B",
	Encoding256b:        "256B/257B (transcoded FEC-enabled data)",
	EncodingPam4:        "PAM4",
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
		"value": e.String(),
		"hex":   hex.EncodeToString([]byte{byte(e)}),
	}
	return json.Marshal(m)
}

func (e *Encoding) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	v := Encoding(b[0])
	e = &v
	return nil
}
