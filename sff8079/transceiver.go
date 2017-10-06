package sff8079

import (
	"encoding/hex"
	"encoding/json"
	"strings"
)

const (
	Ether10gBaseEr    = (1 << 7)
	Ether10gBaseLrm   = (1 << 6)
	Ether10gBaeLr     = (1 << 5)
	Ether10gBaseSr    = (1 << 4)
	Infini1xSx        = (1 << 3)
	Infini1xLx        = (1 << 2)
	Infini1xCopprActv = (1 << 1)
	Infini1xCopprPasv = (1 << 0)

	EsconMmf        = (1 << 7)
	EsconSmf        = (1 << 6)
	SonetOc192Short = (1 << 5)
	SonetReachBit1  = (1 << 4)
	SonetReachBit2  = (1 << 3)
	SonetOc48Long   = (1 << 2)
	SonetOc48Inter  = (1 << 1)
	SonetOc48Short  = (1 << 0)

	SonetOc12Long  = (1 << 6)
	SonetOc12Inter = (1 << 5)
	SonetOc12Short = (1 << 4)
	SonetOc3Long   = (1 << 2)
	SonetOc3Inter  = (1 << 1)
	SonetOc3Short  = (1 << 0)
)

var etherComplCodeNames = map[byte]string{
	Ether10gBaseEr:    "10G Ethernet: 10G Base-ER [SFF-8472 rev10.4 only]",
	Ether10gBaseLrm:   "10G Ethernet: 10G Base-LRM",
	Ether10gBaeLr:     "10G Ethernet: 10G Base-LR",
	Ether10gBaseSr:    "10G Ethernet: 10G Base-SR",
	Infini1xSx:        "Infiniband: 1X SX",
	Infini1xLx:        "Infiniband: 1X LX",
	Infini1xCopprActv: "Infiniband: 1X Copper Active",
	Infini1xCopprPasv: "Infiniband: 1X Copper Passive",
}

var esconSonetCodeNames = map[byte]string{
	EsconMmf:        "ESCON: ESCON MMF, 1310nm LED",
	EsconSmf:        "ESCON: ESCON SMF, 1310nm Laser",
	SonetOc192Short: "SONET: OC-192, short reach",
	SonetReachBit1:  "SONET: SONET reach specifier bit 1",
	SonetReachBit2:  "SONET: SONET reach specifier bit 2",
	SonetOc48Long:   "SONET: OC-48, long reach",
	SonetOc48Inter:  "SONET: OC-48, intermediate reach",
	SonetOc48Short:  "SONET: OC-48, short reach",
}

var sonetCodeNames = map[byte]string{
	SonetOc12Long:  "SONET: OC-12, single mode, long reach",
	SonetOc12Inter: "SONET: OC-12, single mode, inter. reach",
	SonetOc12Short: "SONET: OC-12, short reach",
	SonetOc3Long:   "SONET: OC-3, single mode, long reach",
	SonetOc3Inter:  "SONET: OC-3, single mode, inter. reach",
	SonetOc3Short:  "SONET: OC-3, short reach",
}

type Transceiver [8]byte

func (t Transceiver) List() []string {
	r := []string{}
	b := byte(t[0])
	for k, v := range etherComplCodeNames {
		if k&b != 0 {
			r = append(r, v)
		}
	}

	b = byte(t[1])
	for k, v := range esconSonetCodeNames {
		if k&b != 0 {
			r = append(r, v)
		}
	}

	b = byte(t[2])
	for k, v := range sonetCodeNames {
		if k&b != 0 {
			r = append(r, v)
		}
	}

	return r
}

func (t Transceiver) String() string {
	return strings.Join(t.List(), "\n")
}

func (t Transceiver) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"names": t.List(),
		"hex":   hex.EncodeToString(t[:8]),
	}
	return json.Marshal(m)
}
