package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

type String2 [2]byte
type String16 [16]byte
type VendorOUI [3]byte
type DateCode [8]byte
type ValueKm byte
type ValueM byte
type ValueMBps byte

func (b String2) String() string {
	return strings.TrimSpace(string(b[0:2]))
}

func (b String2) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": b.String(),
		"hex":  hex.EncodeToString([]byte(b[:2])),
	}
	return json.Marshal(m)
}

func (b String16) String() string {
	return strings.TrimSpace(string(b[:16]))
}

func (b String16) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": b.String(),
		"hex":  hex.EncodeToString([]byte(b[:16])),
	}
	return json.Marshal(m)
}

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

func (v ValueKm) String() string {
	return fmt.Sprintf("%d Km", v)
}

func (v ValueKm) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": uint8(v),
		"unit":  "Km",
		"hex":   hex.EncodeToString([]byte{byte(v)}),
	}
	return json.Marshal(m)
}

func (v ValueM) String() string {
	return fmt.Sprintf("%d m", v)
}

func (v ValueM) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": uint8(v),
		"unit":  "m",
		"hex":   hex.EncodeToString([]byte{byte(v)}),
	}
	return json.Marshal(m)
}

func (v ValueMBps) String() string {
	return fmt.Sprintf("%d m", v)
}

func (v ValueMBps) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": uint8(v),
		"unit":  "MBps",
		"hex":   hex.EncodeToString([]byte{byte(v)}),
	}
	return json.Marshal(m)
}

type SFF8636 struct {
	ExtIdentifier     ExtIdentifier `json:"extIdentifier"`     // 129 - Ext. Identifier
	ConnectorType     ConnectorType `json:"connectorType"`     // 130 - Connector Type
	SpecComp          SpecComp      `json:"specComp"`          // 131-138 - Specification Compliance
	Encoding          Encoding      `json:"encoding"`          // 139 - Encoding
	BrNominal         ValueMBps     `json:"brNominal"`         // 140 - BR, nominal
	ExtRateSelComp    byte          `json:"extRateSelComp"`    // 141 - Extended Rate Select Compliance
	LengthSmf         ValueKm       `json:"lengthSmf"`         // 142 - Length (SMF)
	LengthOm3         ValueM        `json:"lengthOm3"`         // 143 - Length (OM3 50 um)
	LengthOm2         ValueM        `json:"lengthOm2"`         // 144 - Length (OM2 50 um)
	LengthOm1         ValueM        `json:"lengthOm1"`         // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopr        ValueM        `json:"lengthCopr"`        // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           byte          `json:"devTech"`           // 147 - Device technology
	Vendor            String16      `json:"vendor"`            // 148-163 - Vendor name
	ExtModule         byte          `json:"extModule"`         // 164 - Extended Module
	VendorOui         VendorOUI     `json:"vendorOui"`         // 165-167 - Vendor OUI
	VendorPn          String16      `json:"vendorPn"`          // 168-183 - Vendor PN
	VendorRev         String2       `json:"vendorRev"`         // 184-185 - Vendor rev
	LaserWavelen      [2]byte       `json:"laserWavelen"`      // 186 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler [2]byte       `json:"laserWavelenToler"` // 187 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte          `json:"maxCaseTempC"`      // 190 - Max case temp.
	CcBase            byte          `json:"ccBase"`            // 191 - CC_BASE
	LinkCodes         LinkCodes     `json:"linkCodes"`         // 192 - Link codes
	Options           [3]byte       `json:"options"`           // 193-195 - Options
	VendorSn          String16      `json:"vendorSn"`          // 196-211 - Vendor SN
	DateCode          DateCode      `json:"dateCode"`          // 212-219 - Date Code
	DiagMonType       byte          `json:"diagMonType"`       // 220 - Diagnostic Monitoring Type
	EnhOptions        byte          `json:"enhOptions"`        // 221 - Enhanced Options
	BrNominalExt      byte          `json:"brNominalExt"`      // 222 - BR, Nominal
	CcExt             byte          `json:"ccExt"`             // 223 - CC_EXT
	VendorSpec        [32]byte      `json:"vendorSpec"`        // 224-255 - Vendor Specific
}

func (s *SFF8636) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *SFF8636) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
