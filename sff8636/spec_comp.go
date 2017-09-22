package sff8636

// INCORRECT!?!

import (
	"encoding/hex"
	"encoding/json"
)

const (
	SpecCompReserved   = (1 << 7)
	SpecComp10GBaseLrm = (1 << 6)
	SpecComp10GBaseLr  = (1 << 5)
	SpecComp10GBaseSr  = (1 << 4)
	SpecComp40GBaseCr4 = (1 << 3)
	SpecComp40GBaseSr4 = (1 << 2)
	SpecComp40GBaseLr4 = (1 << 1)
	SpecComp40GActive  = (1 << 0)
)

var specCompNames = map[byte]string{
	SpecCompReserved:   "10G Ethernet: 10G Base-LRM",
	SpecComp10GBaseLrm: "10G Ethernet: 10G Base-LR",
	SpecComp10GBaseLr:  "10G Ethernet: 10G Base-SR",
	SpecComp10GBaseSr:  "40G Ethernet: 40G Base-CR4",
	SpecComp40GBaseCr4: "40G Ethernet: 40G Base-SR4",
	SpecComp40GBaseSr4: "40G Ethernet: 40G Base-LR4",
	SpecComp40GActive:  "40G Ethernet: 40G Active Cable (XLPPI)",
}

type SpecComp [8]byte

func (s SpecComp) String() string {
	r := ""
	for k, v := range specCompNames {
		b := byte(s[0])
		if k&b != 0 {
			r += v + "\n"
		}
	}
	return r
}

func (s SpecComp) MarshalJSON() ([]byte, error) {
	r := []string{}
	for k, v := range specCompNames {
		b := byte(s[0])
		if k&b != 0 {
			r = append(r, v)
		}
	}

	m := map[string]interface{}{
		"names": r,
		"hex":   hex.EncodeToString([]byte(s[:8])),
	}
	return json.Marshal(m)
}
