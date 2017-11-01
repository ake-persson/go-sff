package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
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

type ExtIdentifier byte

func (e ExtIdentifier) List() []string {
	b := byte(e)
	s := []string{
		pwrClassNames[b&PwrClassMask],
		clieCodeNames[b&ClieCodeMask],
		fmt.Sprintf("%s, %s", cdrInTxNames[b&CdrInTxMask], cdrInRxNames[b&CdrInRxMask]),
	}

	if b&ExtPwrClassMask != ExtPwrClassUnused {
		s = append(s, extPwrClassNames[b&ExtPwrClassMask])
	}

	return s
}

func (e ExtIdentifier) String() string {
	return strings.Join(e.List(), "\n")
}

func (e ExtIdentifier) MarshalJSON() ([]byte, error) {
	b := byte(e)
	m := map[string]interface{}{
		"values": e.List(),
		"hex":    hex.EncodeToString([]byte{b}),
	}
	return json.Marshal(m)
}

func (e *ExtIdentifier) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*e = ExtIdentifier(b[0])
	return nil
}
