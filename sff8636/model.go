package sff8636

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/mickep76/go-sff/common"
)

type Sff8636 struct {
	Identifier        common.Identifier    `json:"identifier"`     // 128 - Identifier
	ExtIdentifier     ExtIdentifier        `json:"extIdentifier"`  // 129 - Ext. Identifier
	ConnectorType     common.ConnectorType `json:"connector"`      // 130 - Connector Type
	Transceiver       Transceiver          `json:"transceiver"`    // 131-138 - Specification Compliance
	Encoding          Encoding             `json:"encoding"`       // 139 - Encoding
	BrNominal         common.Value100Mbps  `json:"brNominal"`      // 140 - BR, nominal
	RateIdentifier    byte                 `json:"rateIdentifier"` // 141 - Extended Rate Select Compliance
	LengthSmf         common.ValueKm       `json:"lengthSmf"`      // 142 - Length (SMF)
	LengthOm3         common.ValueM        `json:"lengthOm3"`      // 143 - Length (OM3 50 um)
	LengthOm2         common.ValueM        `json:"lengthOm2"`      // 144 - Length (OM2 50 um)
	LengthOm1         common.ValueM        `json:"lengthOm1"`      // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopper      common.ValueM        `json:"lengthCopper"`   // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           byte                 `json:"-"`              // 147 - Device technology
	Vendor            common.String16      `json:"vendor"`         // 148-163 - Vendor name
	ExtModule         byte                 `json:"-"`              // 164 - Extended Module
	VendorOui         common.VendorOUI     `json:"vendorOui"`      // 165-167 - Vendor OUI
	VendorPn          common.String16      `json:"vendorPn"`       // 168-183 - Vendor PN
	VendorRev         common.String2       `json:"vendorRev"`      // 184-185 - Vendor rev
	LaserWavelen      [2]byte              `json:"-"`              // 186 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler [2]byte              `json:"-"`              // 187 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte                 `json:"-"`              // 190 - Max case temp.
	CcBase            byte                 `json:"-"`              // 191 - CC_BASE
	LinkCodes         LinkCodes            `json:"linkCodes"`      // 192 - Link codes
	Options           [3]byte              `json:"options"`        // 193-195 - Options
	VendorSn          common.String16      `json:"vendorSn"`       // 196-211 - Vendor SN
	DateCode          common.DateCode      `json:"dateCode"`       // 212-219 - Date Code
	DiagMonType       byte                 `json:"-"`              // 220 - Diagnostic Monitoring Type
	EnhOptions        byte                 `json:"-"`              // 221 - Enhanced Options
	BrNominalExt      byte                 `json:"-"`              // 222 - BR, Nominal
	CcExt             byte                 `json:"-"`              // 223 - CC_EXT
	VendorSpec        [32]byte             `json:"-"`              // 224-255 - Vendor Specific
}

func New(eeprom []byte) (*Sff8636, error) {
	if len(eeprom) != 640 {
		return nil, fmt.Errorf("incorrect size of eeprom for SFF-8636, should be 640 got: %d", len(eeprom))
	}

	return (*Sff8636)(unsafe.Pointer(&eeprom[128])), nil
}

func (s *Sff8636) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *Sff8636) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}

func (s *Sff8636) String() string {
	return fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Identifier [128]", byte(s.Identifier), s.Identifier) +
		fmt.Sprintf("%-50s : 0x%02x\n", "Extended Identifier [129]", byte(s.ExtIdentifier)) +
		fmt.Sprintf("%-50s : %s\n", "Extended Identifier Description", strings.Join(s.ExtIdentifier.List(), fmt.Sprintf("\n%-50s : ", " "))) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Connector [130]", byte(s.ConnectorType), s.ConnectorType) +
		fmt.Sprintf("%-50s : 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x\n", "Transceiver Codes [131-138]", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]) +
		fmt.Sprintf("%-50s : %s\n", "Transceiver Type", strings.Join(s.Transceiver.List(), fmt.Sprintf("\n%-50s : ", " "))) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Encoding [139]", byte(s.Encoding), s.Encoding) +
		fmt.Sprintf("%-50s : %s\n", "BR, Nominal [140]", s.BrNominal) +
		fmt.Sprintf("%-50s : 0x%02x\n", "Rate Identifier [141]", s.RateIdentifier) +
		fmt.Sprintf("%-50s : %s\n", "Length (SMF) [142]", s.LengthSmf) +
		fmt.Sprintf("%-50s : %s\n", "Length (OM3 50um) [143]", s.LengthOm3) +
		fmt.Sprintf("%-50s : %s\n", "Length (OM2 50um) [144]", s.LengthOm2) +
		fmt.Sprintf("%-50s : %s\n", "Length (OM1 62.5um) [145]", s.LengthOm1) +
		fmt.Sprintf("%-50s : %s\n", "Length (Copper or Active cable) [146]", s.LengthCopper) +
		fmt.Sprintf("%-50s : %s\n", "Vendor [148-163]", s.Vendor) +
		fmt.Sprintf("%-50s : %s\n", "Vendor OUI [165-167]", s.VendorOui) +
		fmt.Sprintf("%-50s : %s\n", "Vendor PN [168-183]", s.VendorPn) +
		fmt.Sprintf("%-50s : %s\n", "Vendor Rev [184-185]", s.VendorRev) +
		fmt.Sprintf("%-50s : %s\n", "Vendor SN [196-211]", s.VendorSn) +
		fmt.Sprintf("%-50s : %s\n", "Date Code [212-219]", s.DateCode)
}
