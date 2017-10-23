package sff8636

import (
	"encoding/hex"
	"encoding/json"
)

const (
	LinkCodesUnspecified    = 0x00
	LinkCodes100gAoc        = 0x01
	LinkCodes100gSr4        = 0x02
	LinkCodes100gLr4        = 0x03
	LinkCodes100gEr4        = 0x04
	LinkCodes100gSr10       = 0x05
	LinkCodes100gCwdm4Fec   = 0x06
	LinkCodes100gPsm4       = 0x07
	LinkCodes100gAcc        = 0x08
	LinkCodes100gCwdm4NoFec = 0x09
	LinkCodes100gRsvd1      = 0x0A
	LinkCodes100gCr4        = 0x0B
	LinkCodes25gCrCaS       = 0x0C
	LinkCodes25gCrCaN       = 0x0D
	LinkCodes40gEr4         = 0x10
	LinkCodes4X10Sr         = 0x11
	LinkCodes40gPsm4        = 0x12
	LinkCodesG959P1t12d1    = 0x13
	LinkCodesG959P1s12d2    = 0x14
	LinkCodesG959P1l12d2    = 0x15
	LinkCodes10gTSfi        = 0x16
	LinkCodes100gClr4       = 0x17
	LinkCodes100gAoc2       = 0x18
	LinkCodes100gAcc2       = 0x19
)

var linkCodesNames = map[byte]string{
	LinkCodesUnspecified:    "Reserved or unknown",
	LinkCodes100gAoc:        "100G LinkCodes: 100G AOC or 25GAUI C2M AOC with worst BER of 5x10^(-5)",
	LinkCodes100gSr4:        "100G LinkCodes: 100G Base-SR4 or 25GBase-SR",
	LinkCodes100gLr4:        "100G LinkCodes: 100G Base-LR4",
	LinkCodes100gEr4:        "100G LinkCodes: 100G Base-ER4",
	LinkCodes100gSr10:       "100G LinkCodes: 100G Base-SR10",
	LinkCodes100gCwdm4Fec:   "100G LinkCodes: 100G CWDM4 MSA with FEC",
	LinkCodes100gPsm4:       "100G LinkCodes: 100G PSM4 Parallel SMF",
	LinkCodes100gAcc:        "100G LinkCodes: 100G ACC or 25GAUI C2M ACC with worst BER of 5x10^(-5)",
	LinkCodes100gCwdm4NoFec: "100G LinkCodes: 100G CWDM4 MSA without FEC",
	LinkCodes100gRsvd1:      "(reserved or unknown)",
	LinkCodes100gCr4:        "100G LinkCodes: 100G Base-CR4 or 25G Base-CR CA-L",
	LinkCodes25gCrCaS:       "25G LinkCodes: 25G Base-CR CA-S",
	LinkCodes25gCrCaN:       "25G LinkCodes: 25G Base-CR CA-N",
	LinkCodes40gEr4:         "40G LinkCodes: 40G Base-ER4",
	LinkCodes4X10Sr:         "4x10G LinkCodes: 10G Base-SR",
	LinkCodes40gPsm4:        "40G LinkCodes: 40G PSM4 Parallel SMF",
	LinkCodesG959P1t12d1:    "LinkCodes: G959.1 profile P1I1-2D1 (10709 MBd, 2km, 1310nm SM)",
	LinkCodesG959P1s12d2:    "LinkCodes: G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550nm SM)",
	LinkCodesG959P1l12d2:    "LinkCodes: G959.1 profile P1L1-2D2 (10709 MBd, 80km, 1550nm SM)",
	LinkCodes10gTSfi:        "10G LinkCodes: 10G Base-T with SFI electrical interface",
	LinkCodes100gClr4:       "100G LinkCodes: 100G CLR4",
	LinkCodes100gAoc2:       "100G LinkCodes: 100G AOC or 25GAUI C2M AOC with worst BER of 10^(-12)",
	LinkCodes100gAcc2:       "100G LinkCodes: 100G ACC or 25GAUI C2M ACC with worst BER of 10^(-12)",
}

type LinkCodes byte

func (l LinkCodes) String() string {
	n, ok := linkCodesNames[byte(l)]
	if !ok {
		return "Reserved or unknown"
	}
	return n
}

func (l LinkCodes) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": l.String(),
		"hex":   hex.EncodeToString([]byte{byte(l)}),
	}
	return json.Marshal(m)
}
