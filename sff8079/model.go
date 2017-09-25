package sff8079

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mickep76/go-sff/common"
	"strings"
	"unsafe"
)

type String4 [4]byte
type String16 [16]byte
type VendorOUI [3]byte
type DateCode [8]byte
type ValueMBps byte

func (b String4) String() string {
	return strings.TrimSpace(string(b[0:2]))
}

func (b String4) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name": b.String(),
		"hex":  hex.EncodeToString([]byte(b[:4])),
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

// SFP/SFP+
type SFF8079 struct {
	Identifier      common.Identifier    `json:"identifier"`      // 0 - Identifier
	ExtIdentifier   byte                 `json:"extIdentifier"`   // 1 - Ext. Identifier
	ConnectorType   common.ConnectorType `json:"connectorType"`   // 2 - Connector
	Transc          [8]byte              `json:"transceiver"`     // 3-10 - Transceiver
	Encoding        Encoding             `json:"encoding"`        // 11 - Encoding
	BrNominal       ValueMBps            `json:"brNominal`        // 12 - BR Nominal
	RateId          byte                 `json:"rateId"`          // 13 - Rate ID
	Length9umKm     byte                 `json:"length9umKm"`     // 14 - Length(9μm) - km - (SMF)?
	Length9umM      byte                 `json:"length9umM"`      // 15 - Length (9μm) - (SMF)?
	Length50umM     byte                 `json:"length50umM"`     // 16 - Length (50μm)
	Length625umM    byte                 `json:"length625umM"`    // 17 - Length (62.5um)
	LengthCopper    byte                 `json:"lengthCopper"`    // 18 - Length (Copper)
	Length50umM2    byte                 `json:"length50umM2"`    // 19 - Length (50μm)
	Vendor          String16             `json:"vendor"`          // 20-35 - Vendor name
	TranscComp      byte                 `json:"transceiverComp"` // 36 - Transciever
	VendorOui       VendorOUI            `json:"vendorOUI"`       // 37-39 - Vendor OUI
	VendorPn        String16             `json:"vendorPn"`        // 40-55 - Vendor PN
	VendorRev       String4              `json:"vendorRev"`       // 56-59 - Vendor rev
	LaserWavelength [2]byte              `json:"laserWavelength"` // 60-61 - Laser wavelength
	Unallocated     byte                 `json:"unallocated"`     // 62 - Unallocated
	CcBase          byte                 `json:"ccBase"`          // 63 - CC_BASE
	Options         [2]byte              `json:"options"`         // 64-65 - Options
	BrMax           byte                 `json:"brMax"`           // 66 - BR, max
	BrMin           byte                 `json:"brMin"`           // 67 - BR, min
	VendorSn        String16             `json:"vendorSn"`        // 68-83 - Vendor SN
	DateCode        DateCode             `json:"dateCode"`        // 84-91 - Date code
	DiagMonitType   byte                 `json:"diagMonitType"`   // 92 - Diagnostic Monitoring Type
	EnhancedOpts    byte                 `json:"enhancedOpts"`    // 93 - Enhanced Options
	Sff8472Comp     byte                 `json:"sff8472Comp"`     // 94 - SFF-8472 Compliance
	CcExt           byte                 `json:"ccExt"`           // 95 - CC_EXT
	VendorSpec      [32]byte             `json:"vendorSpec"`      // 96-127 - Vendor Specific
	Reserved        [128]byte            `json:"reserved"`        // 128-255 - Reserved
}

func New(eeprom []byte) (*SFF8079, error) {
	if len(eeprom) != 256 {
		return nil, fmt.Errorf("incorrect size of eeprom for SFF-8079, should be 256 got: %d", len(eeprom))
	}

	return (*SFF8079)(unsafe.Pointer(&eeprom[0])), nil
}

func (s *SFF8079) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *SFF8079) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
