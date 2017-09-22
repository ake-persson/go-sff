package sff8636

import (
	"encoding/hex"
	"encoding/json"
)

const (
	LinkCodesUnspecified    = 0x00
	LinkCodes100GAOC        = 0x01
	LinkCodes100GSR4        = 0x02
	LinkCodes100GLR4        = 0x03
	LinkCodes100GER4        = 0x04
	LinkCodes100GSR10       = 0x05
	LinkCodes100GCWDM4FEC   = 0x06
	LinkCodes100GPSM4       = 0x07
	LinkCodes100GACC        = 0x08
	LinkCodes100GCWDM4NoFEC = 0x09
	LinkCodes100GRSVD1      = 0x0A
	LinkCodes100GCR4        = 0x0B
	LinkCodes25GCRCAS       = 0x0C
	LinkCodes25GCRCAN       = 0x0D
	LinkCodes40GER4         = 0x10
	LinkCodes4X10SR         = 0x11
	LinkCodes40GPSM4        = 0x12
	LinkCodesG959P1I12D1    = 0x13
	LinkCodesG959P1S12D2    = 0x14
	LinkCodesG959P1L12D2    = 0x15
	LinkCodes10GTSFI        = 0x16
	LinkCodes100GCLR4       = 0x17
	LinkCodes100GAOC2       = 0x18
	LinkCodes100GACC2       = 0x19
)

var linkCodesNames = map[byte]string{
	LinkCodesUnspecified:    "Reserved or unknown",
	LinkCodes100GAOC:        "100G LinkCodes: 100G AOC or 25GAUI C2M AOC with worst BER of 5x10^(-5)",
	LinkCodes100GSR4:        "100G LinkCodes: 100G Base-SR4 or 25GBase-SR",
	LinkCodes100GLR4:        "100G LinkCodes: 100G Base-LR4",
	LinkCodes100GER4:        "100G LinkCodes: 100G Base-ER4",
	LinkCodes100GSR10:       "100G LinkCodes: 100G Base-SR10",
	LinkCodes100GCWDM4FEC:   "100G LinkCodes: 100G CWDM4 MSA with FEC",
	LinkCodes100GPSM4:       "100G LinkCodes: 100G PSM4 Parallel SMF",
	LinkCodes100GACC:        "100G LinkCodes: 100G ACC or 25GAUI C2M ACC with worst BER of 5x10^(-5)",
	LinkCodes100GCWDM4NoFEC: "100G LinkCodes: 100G CWDM4 MSA without FEC",
	LinkCodes100GRSVD1:      "(reserved or unknown)",
	LinkCodes100GCR4:        "100G LinkCodes: 100G Base-CR4 or 25G Base-CR CA-L",
	LinkCodes25GCRCAS:       "25G LinkCodes: 25G Base-CR CA-S",
	LinkCodes25GCRCAN:       "25G LinkCodes: 25G Base-CR CA-N",
	LinkCodes40GER4:         "40G LinkCodes: 40G Base-ER4",
	LinkCodes4X10SR:         "4x10G LinkCodes: 10G Base-SR",
	LinkCodes40GPSM4:        "40G LinkCodes: 40G PSM4 Parallel SMF",
	LinkCodesG959P1I12D1:    "LinkCodes: G959.1 profile P1I1-2D1 (10709 MBd, 2km, 1310nm SM)",
	LinkCodesG959P1S12D2:    "LinkCodes: G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550nm SM)",
	LinkCodesG959P1L12D2:    "LinkCodes: G959.1 profile P1L1-2D2 (10709 MBd, 80km, 1550nm SM)",
	LinkCodes10GTSFI:        "10G LinkCodes: 10G Base-T with SFI electrical interface",
	LinkCodes100GCLR4:       "100G LinkCodes: 100G CLR4",
	LinkCodes100GAOC2:       "100G LinkCodes: 100G AOC or 25GAUI C2M AOC with worst BER of 10^(-12)",
	LinkCodes100GACC2:       "100G LinkCodes: 100G ACC or 25GAUI C2M ACC with worst BER of 10^(-12)",
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
		"name": l.String(),
		"hex":  hex.EncodeToString([]byte{byte(l)}),
	}
	return json.Marshal(m)
}
