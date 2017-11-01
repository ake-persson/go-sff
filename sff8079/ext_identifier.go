package sff8079

import (
	"encoding/hex"
	"encoding/json"
)

const (
	ExtIdentifierGbicNotSpec     = 0x00
	ExtIdentifier2WireInterfId   = 0x04
	ExtIdentifierGbicComplModDef = 0x07
)

var extIdentifierNames = map[byte]string{
	ExtIdentifierGbicNotSpec:     "GBIC not specified / not MOD_DEF compliant",
	ExtIdentifier2WireInterfId:   "GBIC/SFP defined by 2-wire interface ID",
	ExtIdentifierGbicComplModDef: "GBIC compliant with MOD_DEF",
}

type ExtIdentifier byte

func (e ExtIdentifier) String() string {
	n, ok := extIdentifierNames[byte(e)]
	if !ok {
		return "Unknown"
	}
	return n
}

func (e ExtIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": e.String(),
		"hex":   hex.EncodeToString([]byte{byte(e)}),
	}
	return json.Marshal(m)
}

func (e *ExtIdentifier) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	v := ExtIdentifier(b[0])
	e = &v
	return nil
}
