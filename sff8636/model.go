package sff8636

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/mickep76/go-sff/common"
)

const (
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
	white   = "\x1b[37m"
	clear   = "\x1b[0m"
)

type Sff8636 struct {
	Identifier        common.Identifier   `json:"identifier"`     // 128 - Identifier
	ExtIdentifier     ExtIdentifier       `json:"extIdentifier"`  // 129 - Ext. Identifier
	Connector         common.Connector    `json:"connector"`      // 130 - Connector Type
	Transceiver       Transceiver         `json:"transceiver"`    // 131-138 - Specification Compliance
	Encoding          Encoding            `json:"encoding"`       // 139 - Encoding
	BrNominal         common.Value100Mbps `json:"brNominal"`      // 140 - BR, nominal
	RateIdentifier    byte                `json:"rateIdentifier"` // 141 - Extended Rate Select Compliance
	LengthSmf         common.ValueKm      `json:"lengthSmf"`      // 142 - Length (SMF)
	LengthOm3         common.ValueM       `json:"lengthOm3"`      // 143 - Length (OM3 50 um)
	LengthOm2         common.ValueM       `json:"lengthOm2"`      // 144 - Length (OM2 50 um)
	LengthOm1         common.ValueM       `json:"lengthOm1"`      // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopper      common.ValueM       `json:"lengthCopper"`   // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           byte                `json:"-"`              // 147 - Device technology
	Vendor            common.String16     `json:"vendor"`         // 148-163 - Vendor name
	ExtModule         byte                `json:"-"`              // 164 - Extended Module
	VendorOui         common.VendorOUI    `json:"vendorOui"`      // 165-167 - Vendor OUI
	VendorPn          common.String16     `json:"vendorPn"`       // 168-183 - Vendor PN
	VendorRev         common.String2      `json:"vendorRev"`      // 184-185 - Vendor rev
	LaserWavelen      [2]byte             `json:"-"`              // 186 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler [2]byte             `json:"-"`              // 187 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte                `json:"-"`              // 190 - Max case temp.
	CcBase            byte                `json:"-"`              // 191 - CC_BASE
	LinkCodes         LinkCodes           `json:"linkCodes"`      // 192 - Link codes
	Options           [3]byte             `json:"options"`        // 193-195 - Options
	VendorSn          common.String16     `json:"vendorSn"`       // 196-211 - Vendor SN
	DateCode          common.DateCode     `json:"dateCode"`       // 212-219 - Date Code
	DiagMonType       byte                `json:"-"`              // 220 - Diagnostic Monitoring Type
	EnhOptions        byte                `json:"-"`              // 221 - Enhanced Options
	BrNominalExt      byte                `json:"-"`              // 222 - BR, Nominal
	CcExt             byte                `json:"-"`              // 223 - CC_EXT
	VendorSpec        [32]byte            `json:"-"`              // 224-255 - Vendor Specific
}

func Decode(eeprom []byte) (*Sff8636, error) {
	if len(eeprom) < 256 {
		return nil, fmt.Errorf("eeprom size to small needs to be 256 bytes or larger got: %d bytes", len(eeprom))
	}

	if eeprom[128] == 12 || eeprom[128] == 13 || eeprom[128] == 17 {
		return (*Sff8636)(unsafe.Pointer(&eeprom[128])), nil
	}

	return nil, fmt.Errorf("unknown eeprom standard, identifier: 0x%02x", byte(eeprom[0]))
}

func (s *Sff8636) String() string {
	return fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Identifier [128]", byte(s.Identifier), s.Identifier) +
		fmt.Sprintf("%-50s : 0x%02x\n", "Extended Identifier [129]", byte(s.ExtIdentifier)) +
		fmt.Sprintf("%-50s : %s\n", "Extended Identifier Description", strings.Join(s.ExtIdentifier.List(), fmt.Sprintf("\n%-50s : ", " "))) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Connector [130]", byte(s.Connector), s.Connector) +
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

func strCol(k string, v string, c1 string, c2 string) string {
	return fmt.Sprintf("%s%-50s%s : %s%s%s\n", c1, k, clear, c2, v, clear)
}

func joinStrCol(k string, l []string, c1 string, c2 string) string {
	r := strCol(k, l[0], c1, c2)
	for _, s := range l[1:] {
		r += strCol("", s, c1, c2)
	}
	return r
}

func (s *Sff8636) StringCol() string {
	return strCol("Identifier [128]", fmt.Sprintf("0x%02x (%s)", byte(s.Identifier), s.Identifier), cyan, green) +
		strCol("Extended Identifier [129]", fmt.Sprintf("0x%02x", byte(s.ExtIdentifier)), cyan, green) +
		strCol("Extended Identifier Description", strings.Join(s.ExtIdentifier.List(), fmt.Sprintf("\n%-50s : ", " ")), cyan, green) +
		strCol("Connector [130]", fmt.Sprintf("0x%02x (%s)", byte(s.Connector), s.Connector), cyan, green) +
		strCol("Transceiver Codes [131-138]", fmt.Sprintf("0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]), cyan, green) +
		joinStrCol("Transceiver Type", s.Transceiver.List(), cyan, yellow) +
		strCol("Encoding [139]", fmt.Sprintf("0x%02x (%s)", byte(s.Encoding), s.Encoding), cyan, green) +
		strCol("BR, Nominal [140]", s.BrNominal.String(), cyan, green) +
		strCol("Rate Identifier [141]", fmt.Sprintf("0x%02x", s.RateIdentifier), cyan, green) +
		strCol("Length (SMF) [142]", s.LengthSmf.String(), cyan, green) +
		strCol("Length (OM3 50um) [143]", s.LengthOm3.String(), cyan, green) +
		strCol("Length (OM2 50um) [144]", s.LengthOm2.String(), cyan, green) +
		strCol("Length (OM1 62.5um) [145]", s.LengthOm1.String(), cyan, green) +
		strCol("Length (Copper or Active cable) [146]", s.LengthCopper.String(), cyan, green) +
		strCol("Vendor [148-163]", s.Vendor.String(), cyan, green) +
		strCol("Vendor OUI [165-167]", s.VendorOui.String(), cyan, green) +
		strCol("Vendor PN [168-183]", s.VendorPn.String(), cyan, green) +
		strCol("Vendor Rev [184-185]", s.VendorRev.String(), cyan, green) +
		strCol("Vendor SN [196-211]", s.VendorSn.String(), cyan, green) +
		strCol("Date Code [212-219]", s.DateCode.String(), cyan, green)
}
