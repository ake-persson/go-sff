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

type Transceiver [8]byte

func (t Transceiver) List() []string {
	r := []string{}
	for k, v := range etherComplCodeNames {
		b := byte(t[0])
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
