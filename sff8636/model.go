package main

import (
	"encoding/json"
)

type SFF8636 struct {
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

func (s *SFF8636) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *SFF8636) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}