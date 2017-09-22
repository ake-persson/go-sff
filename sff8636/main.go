package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	PwrClassMask = 0xC0
	PwrClass1    = (0 << 6)
	PwrClass2    = (1 << 6)
	PwrClass3    = (2 << 6)
	PwrClass4    = (3 << 6)

	ClieCodeMask = 0x10
	NoClieCode   = (0 << 4)
	ClieCode     = (1 << 4)

	CdrInTxMask = 0x08
	NoCdrInTx   = (0 << 3)
	CdrInTx     = (1 << 3)

	CdrInRxMask = 0x04
	NoCdrInRx   = (0 << 2)
	CdrInRx     = (1 << 2)

	ExtPwrClassMask   = 0x03
	ExtPwrClassUnused = 0
	ExtPwrClass5      = 1
	ExtPwrClass6      = 2
	ExtPwrClass7      = 3
)

var pwrClassNames = map[byte]string{
	PwrClass1: "1.5 W max. power consumption",
	PwrClass2: "2.0 W max. power consumption",
	PwrClass3: "2.5 W max. power consumption",
	PwrClass4: "3.5 W max. power consumption",
}

var clieCodeNames = map[byte]string{
	NoClieCode: "No CLEI code present",
	ClieCode:   "CLEI code present",
}

var cdrInTxNames = map[byte]string{
	NoCdrInTx: "No CDR in TX",
	CdrInTx:   "CDR in TX",
}

var cdrInRxNames = map[byte]string{
	NoCdrInRx: "No CDR in RX",
	CdrInRx:   "CDR in RX",
}

var extPwrClassNames = map[byte]string{
	ExtPwrClassUnused: "unused (legacy setting)",
	ExtPwrClass5:      "4.0 W max. power consumption",
	ExtPwrClass6:      "4.5 W max. power consumption",
	ExtPwrClass7:      "5.0 W max. power consumption",
}

type PwrMode byte
type ExtIdentifier byte
type ByteString2 [2]byte
type ByteString16 [16]byte
type VendorOUI [3]byte
type DateCode [8]byte

func (e ExtIdentifier) String() string {
	b := byte(e)
	s := pwrClassNames[b&PwrClassMask] + "\n"
	s += clieCodeNames[b&ClieCodeMask] + "\n"
	s += cdrInTxNames[b&CdrInTxMask] + ", " + cdrInRxNames[b&CdrInRxMask] + "\n"

	if b&ExtPwrClassMask != ExtPwrClassUnused {
		s += extPwrClassNames[b&ExtPwrClassMask] + "\n"
	}
	return s
}

func (b ByteString2) String() string {
	return strings.TrimSpace(string(b[0:2]))
}

func (b ByteString16) String() string {
	return strings.TrimSpace(string(b[0:16]))
}

func (v VendorOUI) String() string {
	return fmt.Sprintf("%x:%x:%x", v[0], v[1], v[2])
}

func (d DateCode) String() string {
	return fmt.Sprintf("20%s-%s-%s", string(d[0:2]), string(d[2:4]), string(d[4:6]))
}

type SFF8636 struct {
	Identifier byte `json:"identifier"` // 0 - Identifier

	padding1 [92]byte

	PwrMode PwrMode `json:"pwrMode"` // 93 - Power Mode

	padding2 [35]byte

	ExtIdentifier     ExtIdentifier `json:"extIdentifier"`     // 129 - Ext. Identifier
	ConnectorType     byte          `json:"connectorType"`     // 130 - Connector Type
	SpecComp          [8]byte       `json:"specComp"`          // 131-138 - Specification Compliance
	Encoding          byte          `json:"encoding"`          // 139 - Encoding
	BrNominal         byte          `json:"brNominal"`         // 140 - BR, nominal
	ExtRateSelComp    byte          `json:"extRateSelComp"`    // 141 - Extended Rate Select Compliance
	LengthSmf         byte          `json:"lengthSmf"`         // 142 - Length (SMF)
	LengthOm3         byte          `json:"lengthOm3"`         // 143 - Length (OM3 50 um)
	LengthOm2         byte          `json:"lengthOm2"`         // 144 - Length (OM2 50 um)
	LengthOm1         byte          `json:"lengthOm1"`         // 145 - Length (OM1 62.5 um) or Copper Cable Attenuation
	LengthCopr        byte          `json:"lengthCopr"`        // 146 - Length (passive copper or active cable or OM4 50 um)
	DevTech           byte          `json:"devTech"`           // 147 - Device technology
	Vendor            ByteString16  `json:"vendor"`            // 148-163 - Vendor name
	ExtModule         byte          `json:"extModule"`         // 164 - Extended Module
	VendorOui         VendorOUI     `json:"vendorOui"`         // 165-167 - Vendor OUI
	VendorPn          ByteString16  `json:"vendorPn"`          // 168-183 - Vendor PN
	VendorRev         ByteString2   `json:"vendorRev"`         // 184-185 - Vendor rev
	LaserWavelen      [2]byte       `json:"laserWavelen"`      // 186 - Wavelength or Copper Cable Attenuation
	LaserWavelenToler [2]byte       `json:"laserWavelenToler"` // 187 - Wavelength tolerance or Copper Cable Attenuation
	MaxCaseTempC      byte          `json:"maxCaseTempC"`      // 190 - Max case temp.
	CcBase            byte          `json:"ccBase"`            // 191 - CC_BASE
	LinkCodes         byte          `json:"linkCodes"`         // 192 - Link codes
	Options           [3]byte       `json:"options"`           // 193-195 - Options
	VendorSn          ByteString16  `json:"vendorSn"`          // 196-211 - Vendor SN
	DateCode          DateCode      `json:"dateCode"`          // 212-219 - Date Code
	DiagMonType       byte          `json:"diagMonType"`       // 220 - Diagnostic Monitoring Type
	EnhOptions        byte          `json:"enhOptions"`        // 221 - Enhanced Options
	BrNominalExt      byte          `json:"brNominalExt"`      // 222 - BR, Nominal
	CcExt             byte          `json:"ccExt"`             // 223 - CC_EXT
	VendorSpec        [32]byte      `json:"vendorSpec"`        // 224-255 - Vendor Specific
}

func main() {
	var b []byte
	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
	} else {
		log.Fatalf("stdin is hungry, feed me")
	}

	// Decode hex
	eeprom, err := hex.DecodeString(strings.TrimRight(string(b), "\n"))
	if err != nil {
		log.Fatalf("decode hex: %v", err)
	}
	fmt.Printf("Eeprom Size: %d\n", len(eeprom))

	if len(eeprom) == 512 || len(eeprom) == 640 {
		s := (*SFF8636)(unsafe.Pointer(&eeprom[0]))
		fmt.Printf("%+v\n", s)
	}
}
