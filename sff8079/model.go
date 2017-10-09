package sff8079

import (
	"encoding/json"
	"fmt"
	"github.com/mickep76/go-sff/common"
	"strings"
	"unsafe"
)

type Sff8079 struct {
	Identifier      common.Identifier    `json:"identifier"`      // 0 - Identifier
	ExtIdentifier   ExtIdentifier        `json:"extIdentifier"`   // 1 - Ext. Identifier
	ConnectorType   common.ConnectorType `json:"connectorType"`   // 2 - Connector
	Transceiver     Transceiver          `json:"transceiver"`     // 3-10 - Transceiver
	Encoding        Encoding             `json:"encoding"`        // 11 - Encoding
	BrNominal       common.ValueMBps     `json:"brNominal`        // 12 - BR Nominal
	RateId          byte                 `json:"rateId"`          // 13 - Rate ID
	LengthSmfKm     common.ValueKm       `json:"lengthSmfKm"`     // 14 - Length(9μm) - km - (SMF)?
	LengthSmfM      common.ValueM        `json:"lengthSmfM"`      // 15 - Length (9μm) - (SMF)?
	Length50umM     common.ValueM        `json:"length50umM"`     // 16 - Length (50μm)
	Length625umM    common.ValueM        `json:"length625umM"`    // 17 - Length (62.5um)
	LengthCopper    common.ValueM        `json:"lengthCopper"`    // 18 - Length (Copper)
	LengthOm3       common.ValueM        `json:"lengthOm3"`       // 19 - Length (50μm)
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

// Use Go text template?
func (s *Sff8079) String() string {
	return "Identifier:\t\t" + s.Identifier.String() +
		"\nExt. Identifier:\t" + s.ExtIdentifier.String() +
		"\nConnector Type:\t\t" + s.ConnectorType.String() +
		"\nTransceiver:\n\t\t\t" + strings.Join(s.Transceiver.List(), "\n\t\t\t") + "\n" +
		"\nEncoding:\t\t" + s.Encoding.String() +
		"\nBR Nominal:\t\t" + s.BrNominal.String() +
		"\nLength (SMF):\t\t" + s.LengthSmfKm.String() +
		"\nLength (SMF):\t\t" + s.LengthSmfM.String() +
		"\nLength (50um):\t\t" + s.Length50umM.String() +
		"\nLength (62.5um):\t" + s.Length625umM.String() +
		"\nLength (Copper):\t" + s.LengthCopper.String() +
		"\nLength (OM3):\t\t" + s.LengthOm3.String() +
		"\nVendor:\t\t" + s.Vendor.String() +
		"\nVendor OUI:\t\t" + s.VendorOui.String() +
		"\nVendor PN:\t\t" + s.VendorPn.String() +
		"\nVendor Rev:\t\t" + s.VendorRev.String()
}

func (s *Sff8079) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *Sff8079) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
