package sff8636

import (
	"encoding/json"
)

const (
	PwrModeHigh     = (1 << 2)
	PwrModeLow      = (1 << 1)
	PwrModeOverride = (1 << 0)
)

var pwrModeNames = map[byte]string{
	PwrModeHigh: "High power class (> 3.5 W) enabled",
	PwrModeLow:  "High power class (> 3.5 W) not enabled",
}

type PwrMode byte

func (p PwrMode) String() string {
	b := byte(p)
	if b&PwrModeHigh != 0 {
		return pwrModeNames[PwrModeHigh]
	}
	return pwrModeNames[PwrModeLow]
}

func (p PwrMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
