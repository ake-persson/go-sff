package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func stringToJSON(b []byte) ([]byte, error) {
	m := map[string]interface{}{
		"name": string(b),
		"hex":  hex.EncodeToString(b),
	}
	return json.Marshal(m)
}

type String2 [2]byte
type String4 [4]byte
type String16 [16]byte

func (s String2) String() string {
	return strings.TrimSpace(string([]byte(s[:2])))
}

func (s String2) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:2]))
}

func (s String4) String() string {
	return strings.TrimSpace(string([]byte(s[:4])))
}

func (s String4) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:4]))
}

func (s String16) String() string {
	return strings.TrimSpace(string([]byte(s[:16])))
}

func (s String16) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:16]))
}

type ValueM byte
type ValueKm byte
type Value100Mbps byte
type ValuePerc byte

func valueToJSON(b byte, unit string) ([]byte, error) {
	m := map[string]interface{}{
		"value": uint8(b),
		"unit":  unit,
		"hex":   hex.EncodeToString([]byte{byte(b)}),
	}
	return json.Marshal(m)
}

func (v ValueM) String() string {
	return fmt.Sprintf("%d m", v)
}

func (v ValueM) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "m")
}

func (v ValueKm) String() string {
	return fmt.Sprintf("%d km", v)
}

func (v ValueKm) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "km")
}

func (v Value100Mbps) String() string {
	return fmt.Sprintf("%d Mb/s", uint(v)*100)
}

func (v Value100Mbps) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": uint(v) * 100,
		"unit":  "Mb/s",
		"hex":   hex.EncodeToString([]byte{byte(v)}),
	}
	return json.Marshal(m)
}

func (v ValuePerc) String() string {
	return fmt.Sprintf("%d %%", v)
}

func (v ValuePerc) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "%")
}

type VendorOUI [3]byte

func (v VendorOUI) String() string {
	return fmt.Sprintf("%x:%x:%x", v[0], v[1], v[2])
}

func (v VendorOUI) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": v.String(),
		"hex":  hex.EncodeToString([]byte(v[:3])),
	}
	return json.Marshal(m)
}

type DateCode [8]byte

func (d DateCode) String() string {
	return fmt.Sprintf("20%s-%s-%s", string(d[:2]), string(d[2:4]), string(d[4:6]))
}

func (d DateCode) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": d.String(),
		"hex":  hex.EncodeToString([]byte(d[:8])),
	}
	return json.Marshal(m)
}
