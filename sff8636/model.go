package sff8636

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/mickep76/go-sff/common"
)

type SFF8636 struct {
	Identifier        common.Identifier    `json:"identifier"`        // 128 - Identifier
	ExtIdentifier     ExtIdentifier        `json:"extIdentifier"`     // 129 - Ext. Identifier
	ConnectorType     common.ConnectorType `json:"connectorType"`     // 130 - Connector Type
	SpecComp          SpecComp             `json:"specComp"`          // 131-138 - Specification Compliance
	Encoding          Encoding             `json:"encoding"`          // 139 - Encoding
	BrNominal         common.ValueMBps     `json:"brNominal"`         // 140 - BR, nominal
	ExtRateSelComp    byte                 `json:"extRateSelComp"`    // 141 - Extended Rate Select Compliance
	LengthSmf         common.ValueKm       `json:"lengthSmf"`         // 142 - Length (SMF)
	LengthOm3         common.ValueM        `json:"lengthOm3"`         // 143 - Length (OM3 50 um)
	LengthOm2         common.ValueM        `json:"lengthOm2"`         // 144 - Length (OM2 50 um)
	LengthOm1         common.ValueM        `json:"lengthOm1"`         // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopr        common.ValueM        `json:"lengthCopr"`        // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           byte                 `json:"devTech"`           // 147 - Device technology
	Vendor            common.String16      `json:"vendor"`            // 148-163 - Vendor name
	ExtModule         byte                 `json:"extModule"`         // 164 - Extended Module
	VendorOui         common.VendorOUI     `json:"vendorOui"`         // 165-167 - Vendor OUI
	VendorPn          common.String16      `json:"vendorPn"`          // 168-183 - Vendor PN
	VendorRev         common.String2       `json:"vendorRev"`         // 184-185 - Vendor rev
	LaserWavelen      [2]byte              `json:"laserWavelen"`      // 186 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler [2]byte              `json:"laserWavelenToler"` // 187 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte                 `json:"maxCaseTempC"`      // 190 - Max case temp.
	CcBase            byte                 `json:"ccBase"`            // 191 - CC_BASE
	LinkCodes         LinkCodes            `json:"linkCodes"`         // 192 - Link codes
	Options           [3]byte              `json:"options"`           // 193-195 - Options
	VendorSn          common.String16      `json:"vendorSn"`          // 196-211 - Vendor SN
	DateCode          common.DateCode      `json:"dateCode"`          // 212-219 - Date Code
	DiagMonType       byte                 `json:"diagMonType"`       // 220 - Diagnostic Monitoring Type
	EnhOptions        byte                 `json:"enhOptions"`        // 221 - Enhanced Options
	BrNominalExt      byte                 `json:"brNominalExt"`      // 222 - BR, Nominal
	CcExt             byte                 `json:"ccExt"`             // 223 - CC_EXT
	VendorSpec        [32]byte             `json:"vendorSpec"`        // 224-255 - Vendor Specific
}

func New(eeprom []byte) (*SFF8636, error) {
	if len(eeprom) != 640 {
		return nil, fmt.Errorf("incorrect size of eeprom for SFF-8636, should be 640 got: %d", len(eeprom))
	}

	return (*SFF8636)(unsafe.Pointer(&eeprom[128])), nil
}

func (s *SFF8636) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *SFF8636) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
