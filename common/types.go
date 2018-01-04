package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func stringToJSON(b []byte) ([]byte, error) {
	m := map[string]interface{}{
		"value": string(b),
		"hex":   hex.EncodeToString(b),
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

func (s *String2) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then String2 type")
	}

	*s = String2{}
	for i := 0; i < 2; i++ {
		s[i] = b[i]
	}
	return nil
}

func (s String4) String() string {
	return strings.TrimSpace(string([]byte(s[:4])))
}

func (s String4) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:4]))
}

func (s *String4) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 4 {
		return fmt.Errorf("length is shorter then String4 type")
	}

	*s = String4{}
	for i := 0; i < 4; i++ {
		s[i] = b[i]
	}
	return nil
}

func (s String16) String() string {
	return strings.TrimSpace(string([]byte(s[:16])))
}

func (s String16) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:16]))
}

func (s *String16) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 16 {
		return fmt.Errorf("length is shorter then String16 type")
	}

	*s = String16{}
	for i := 0; i < 16; i++ {
		s[i] = b[i]
	}
	return nil
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

func (v *ValueM) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValueM(b[0])
	return nil
}

func (v ValueKm) String() string {
	return fmt.Sprintf("%d km", v)
}

func (v ValueKm) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "km")
}

func (v *ValueKm) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValueKm(b[0])
	return nil
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

func (v *Value100Mbps) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = Value100Mbps(b[0])
	return nil
}

func (v ValuePerc) String() string {
	return fmt.Sprintf("%d %%", v)
}

func (v ValuePerc) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "%")
}

func (v *ValuePerc) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValuePerc(b[0])
	return nil
}

type VendorOUI [3]byte

func (v VendorOUI) String() string {
	return fmt.Sprintf("%x:%x:%x", v[0], v[1], v[2])
}

func (v VendorOUI) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": v.String(),
		"hex":   hex.EncodeToString([]byte(v[:3])),
	}
	return json.Marshal(m)
}

func (v *VendorOUI) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 3 {
		return fmt.Errorf("length is shorter then VendorOUI type")
	}

	*v = VendorOUI{}
	for i := 0; i < 3; i++ {
		v[i] = b[i]
	}
	return nil
}

type DateCode [8]byte

func (d DateCode) String() string {
	return fmt.Sprintf("20%s-%s-%s", string(d[:2]), string(d[2:4]), string(d[4:6]))
}

func (d DateCode) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": d.String(),
		"hex":   hex.EncodeToString([]byte(d[:8])),
	}
	return json.Marshal(m)
}

func (d *DateCode) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 8 {
		return fmt.Errorf("length is shorter then DateCode type")
	}

	*d = DateCode{}
	for i := 0; i < 8; i++ {
		d[i] = b[i]
	}
	return nil
}
