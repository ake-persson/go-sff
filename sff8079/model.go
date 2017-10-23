package sff8079

import (
	"encoding/json"
	"fmt"
	"github.com/mickep76/go-sff/common"
	"strings"
	"unsafe"
)

type Sff8079 struct {
	Identifier      common.Identifier   `json:"identifier"`     // 0 - Identifier
	ExtIdentifier   ExtIdentifier       `json:"extIdentifier"`  // 1 - Ext. Identifier
	Connector       common.Connector    `json:"connector"`      // 2 - Connector
	Transceiver     Transceiver         `json:"transceiver"`    // 3-10 - Transceiver
	Encoding        Encoding            `json:"encoding"`       // 11 - Encoding
	BrNominal       common.Value100Mbps `json:"brNominal`       // 12 - BR Nominal
	RateIdentifier  byte                `json:"rateIdentifier"` // 13 - Rate ID
	LengthSmfKm     common.ValueKm      `json:"lengthSmfKm"`    // 14 - Length(9μm) - km - (SMF)?
	LengthSmfM      common.ValueM       `json:"lengthSmfM"`     // 15 - Length (9μm) - (SMF)?
	Length50umM     common.ValueM       `json:"length50umM"`    // 16 - Length (50μm)
	Length625umM    common.ValueM       `json:"length625umM"`   // 17 - Length (62.5um)
	LengthCopper    common.ValueM       `json:"lengthCopper"`   // 18 - Length (Copper)
	LengthOm3       common.ValueM       `json:"lengthOm3"`      // 19 - Length (50μm)
	Vendor          common.String16     `json:"vendor"`         // 20-35 - Vendor name
	TranscComp      byte                `json:"-"`              // 36 - Transciever
	VendorOui       common.VendorOUI    `json:"vendorOUI"`      // 37-39 - Vendor OUI
	VendorPn        common.String16     `json:"vendorPn"`       // 40-55 - Vendor PN
	VendorRev       common.String4      `json:"vendorRev"`      // 56-59 - Vendor rev
	LaserWavelength [2]byte             `json:"-"`              // 60-61 - Laser wavelength
	Unallocated     byte                `json:"-"`              // 62 - Unallocated
	CcBase          byte                `json:"-"`              // 63 - CC_BASE
	Options         [2]byte             `json:"options"`        // 64-65 - Options
	BrMax           common.ValuePerc    `json:"brMax"`          // 66 - BR, max
	BrMin           common.ValuePerc    `json:"brMin"`          // 67 - BR, min
	VendorSn        common.String16     `json:"vendorSn"`       // 68-83 - Vendor SN
	DateCode        common.DateCode     `json:"dateCode"`       // 84-91 - Date code
	DiagMonitType   byte                `json:"-"`              // 92 - Diagnostic Monitoring Type
	EnhancedOpts    byte                `json:"-"`              // 93 - Enhanced Options
	Sff8472Comp     byte                `json:"-"`              // 94 - SFF-8472 Compliance
	CcExt           byte                `json:"-"`              // 95 - CC_EXT
	VendorSpec      [32]byte            `json:"-"`              // 96-127 - Vendor Specific
	Reserved        [128]byte           `json:"-"`              // 128-255 - Reserved
}

func New(eeprom []byte) (*Sff8079, error) {
	if len(eeprom) != 256 {
		return nil, fmt.Errorf("incorrect size of eeprom for SFF-8079, should be 256 got: %d", len(eeprom))
	}

	return (*Sff8079)(unsafe.Pointer(&eeprom[0])), nil
}

func (s *Sff8079) String() string {
	return fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Identifier [0]", byte(s.Identifier), s.Identifier) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Extended Identifier [1]", byte(s.ExtIdentifier), s.ExtIdentifier) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Connector [2]", byte(s.Connector), s.Connector) +
		fmt.Sprintf("%-50s : 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x\n", "Transceiver Codes [3-10]", s.Transceiver[0], s.Transceiver[1], s.Transceiver[2], s.Transceiver[3], s.Transceiver[4], s.Transceiver[5], s.Transceiver[6], s.Transceiver[7]) +
		fmt.Sprintf("%-50s : %s\n", "Transceiver Type", strings.Join(s.Transceiver.List(), fmt.Sprintf("\n%-50s : ", " "))) +
		fmt.Sprintf("%-50s : 0x%02x (%s)\n", "Encoding [11]", byte(s.Encoding), s.Encoding) +
		fmt.Sprintf("%-50s : %s\n", "BR, Nominal [12]", s.BrNominal) +
		fmt.Sprintf("%-50s : 0x%02x\n", "Rate Identifier [13]", s.RateIdentifier) +
		fmt.Sprintf("%-50s : %s\n", "Length (SMF) [14]", s.LengthSmfKm.String()) +
		fmt.Sprintf("%-50s : %s\n", "Length (SMF) [15]", s.LengthSmfM) +
		fmt.Sprintf("%-50s : %s\n", "Length (50um) [16]", s.Length50umM) +
		fmt.Sprintf("%-50s : %s\n", "Length (62.5um) [17]", s.Length625umM) +
		fmt.Sprintf("%-50s : %s\n", "Length (Copper) [18]", s.LengthCopper) +
		fmt.Sprintf("%-50s : %s\n", "Length (OM3) [19]", s.LengthOm3) +
		fmt.Sprintf("%-50s : %s\n", "Vendor [20-35]", s.Vendor) +
		fmt.Sprintf("%-50s : %s\n", "Vendor OUI [37-39]", s.VendorOui) +
		fmt.Sprintf("%-50s : %s\n", "Vendor PN [40-55]", s.VendorPn) +
		fmt.Sprintf("%-50s : %s\n", "Vendor Rev [56-59]", s.VendorRev) +
		fmt.Sprintf("%-50s : 0x%02x 0x%02x\n", "Option Values [64-65]", s.Options[0], s.Options[1]) +
		fmt.Sprintf("%-50s : %s\n", "BR Margin, Max [66]", s.BrMax) +
		fmt.Sprintf("%-50s : %s\n", "BR Margin, Min [67]", s.BrMin) +
		fmt.Sprintf("%-50s : %s\n", "Vendor SN [68-83]", s.VendorSn) +
		fmt.Sprintf("%-50s : %s\n", "Date Code [84-91]", s.DateCode)
}

func (s *Sff8079) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *Sff8079) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
