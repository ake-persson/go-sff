package sff8079

import (
	"encoding/json"
	"fmt"
	"github.com/mickep76/go-sff/common"
	"unsafe"
)

type Sff8079 struct {
	Identifier      common.Identifier    `json:"identifier"`      // 0 - Identifier
	ExtIdentifier   ExtIdentifier        `json:"extIdentifier"`   // 1 - Ext. Identifier
	ConnectorType   common.ConnectorType `json:"connectorType"`   // 2 - Connector
	Transc          Transceiver          `json:"transceiver"`     // 3-10 - Transceiver
	Encoding        Encoding             `json:"encoding"`        // 11 - Encoding
	BrNominal       common.ValueMBps     `json:"brNominal`        // 12 - BR Nominal
	RateId          byte                 `json:"rateId"`          // 13 - Rate ID
	Length9umKm     common.ValueKm       `json:"length9umKm"`     // 14 - Length(9μm) - km - (SMF)?
	Length9umM      common.ValueM        `json:"length9umM"`      // 15 - Length (9μm) - (SMF)?
	Length50umM     common.ValueM        `json:"length50umM"`     // 16 - Length (50μm)
	Length625umM    common.ValueM        `json:"length625umM"`    // 17 - Length (62.5um)
	LengthCopper    common.ValueM        `json:"lengthCopper"`    // 18 - Length (Copper)
	Length50umM2    common.ValueM        `json:"length50umM2"`    // 19 - Length (50μm)
	Vendor          common.String16      `json:"vendor"`          // 20-35 - Vendor name
	TranscComp      byte                 `json:"transceiverComp"` // 36 - Transciever
	VendorOui       common.VendorOUI     `json:"vendorOUI"`       // 37-39 - Vendor OUI
	VendorPn        common.String16      `json:"vendorPn"`        // 40-55 - Vendor PN
	VendorRev       common.String4       `json:"vendorRev"`       // 56-59 - Vendor rev
	LaserWavelength [2]byte              `json:"laserWavelength"` // 60-61 - Laser wavelength
	Unallocated     byte                 `json:"unallocated"`     // 62 - Unallocated
	CcBase          byte                 `json:"ccBase"`          // 63 - CC_BASE
	Options         [2]byte              `json:"options"`         // 64-65 - Options
	BrMax           byte                 `json:"brMax"`           // 66 - BR, max
	BrMin           byte                 `json:"brMin"`           // 67 - BR, min
	VendorSn        common.String16      `json:"vendorSn"`        // 68-83 - Vendor SN
	DateCode        common.DateCode      `json:"dateCode"`        // 84-91 - Date code
	DiagMonitType   byte                 `json:"diagMonitType"`   // 92 - Diagnostic Monitoring Type
	EnhancedOpts    byte                 `json:"enhancedOpts"`    // 93 - Enhanced Options
	Sff8472Comp     byte                 `json:"sff8472Comp"`     // 94 - SFF-8472 Compliance
	CcExt           byte                 `json:"ccExt"`           // 95 - CC_EXT
	VendorSpec      [32]byte             `json:"vendorSpec"`      // 96-127 - Vendor Specific
	Reserved        [128]byte            `json:"reserved"`        // 128-255 - Reserved
}

func New(eeprom []byte) (*Sff8079, error) {
	if len(eeprom) != 256 {
		return nil, fmt.Errorf("incorrect size of eeprom for SFF-8079, should be 256 got: %d", len(eeprom))
	}

	return (*Sff8079)(unsafe.Pointer(&eeprom[0])), nil
}

func (s *Sff8079) String() string {
	return ""
}

func (s *Sff8079) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *Sff8079) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
