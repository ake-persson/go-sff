package sff8636

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

const (
	Ethernet10gLrm    = (1 << 6)
	Ethernet10gLr     = (1 << 5)
	Ethernet10gSr     = (1 << 4)
	Ethernet40gCr4    = (1 << 3)
	Ethernet40gSr4    = (1 << 2)
	Ethernet40gLr4    = (1 << 1)
	Ethernet40gActive = (1 << 0)
	Sonet40gOtn       = (1 << (3 + 8))
	SonetOc48Lr       = (1 << (2 + 8))
	SonetOc48Ir       = (1 << (1 + 8))
	SonetOc48Sr       = (1 << (0 + 8))
	Sas6g             = (1 << (5 + 16))
	Sas3g             = (1 << (4 + 16))
	Gige1000BaseT     = (1 << (3 + 24))
	Gige1000BaseCx    = (1 << (2 + 24))
	Gige1000BaseLx    = (1 << (1 + 24))
	Gige1000BaseSx    = (1 << (0 + 24))
	FcLenVeryLong     = (1 << (7 + 32))
	FcLenShort        = (1 << (6 + 32))
	FcLenInt          = (1 << (5 + 32))
	FcLenLong         = (1 << (4 + 32))
	FcLenMed          = (1 << (3 + 32))
	FcTechLongLc      = (1 << (1 + 32))
	FcTechElecInter   = (1 << (0 + 32))
	FcTechElecIntra   = (1 << (7 + 40))
	FcTechShortWoOfc  = (1 << (6 + 40))
	FcTechShortWOfc   = (1 << (5 + 40))
	FcTechLongLl      = (1 << (4 + 40))
	FcTransMeidaTw    = (1 << (7 + 48))
	FcTransMediaTp    = (1 << (6 + 48))
	FcTransMediaMi    = (1 << (5 + 48))
	FcTransMediaTv    = (1 << (4 + 48))
	FcTransMediaM6    = (1 << (3 + 48))
	FcTransMediaM5    = (1 << (2 + 48))
	FcTransMediaOm3   = (1 << (1 + 48))
	FcTransMediaSm    = (1 << (0 + 48))
	FcSpeed1200Mbps   = (1 << (7 + 56))
	FcSpeed800Mbps    = (1 << (6 + 56))
	FcSpeed1600Mbps   = (1 << (5 + 56))
	FcSpeed400Mbps    = (1 << (4 + 56))
	FcSpeed200Mbps    = (1 << (2 + 56))
	FcSpeed100Mbps    = (1 << (0 + 56))

	EthernetUnspecified    = 0x00
	Ethernet100gAoc        = 0x01
	Ethernet100gSr4        = 0x02
	Ethernet100gLr4        = 0x03
	Ethernet100gEr4        = 0x04
	Ethernet100gSr10       = 0x05
	Ethernet100gCwdm4Fec   = 0x06
	Ethernet100gPsm4       = 0x07
	Ethernet100gAcc        = 0x08
	Ethernet100gCwdm4NoFec = 0x09
	Ethernet100gRsvd1      = 0x0A
	Ethernet100gCr4        = 0x0B
	Ethernet25gCrCaS       = 0x0C
	Ethernet25gCrCaN       = 0x0D
	Ethernet40gEr4         = 0x10
	Ethernet4X10Sr         = 0x11
	Ethernet40gPsm4        = 0x12
	EthernetG959P1i1_2d1   = 0x13
	EthernetG959P1s1_2d2   = 0x14
	EthernetG959P1l1_2d2   = 0x15
	Ethernet10GtSfi        = 0x16
	Ethernet100gClr4       = 0x17
	Ethernet100gAoc2       = 0x18
	Ethernet100gAcc2       = 0x19
)

var extNames = map[byte]string{
	EthernetUnspecified:    "Reserved or unknown",
	Ethernet100gAoc:        "100G Ethernet: 100G AOC or 25GAUI C2M AOC with worst BER of 5x10^(-5)",
	Ethernet100gSr4:        "100G Ethernet: 100G Base-SR4 or 25GBase-SR",
	Ethernet100gLr4:        "100G Ethernet: 100G Base-LR4",
	Ethernet100gEr4:        "100G Ethernet: 100G Base-ER4",
	Ethernet100gSr10:       "100G Ethernet: 100G Base-SR10",
	Ethernet100gCwdm4Fec:   "100G Ethernet: 100G CWDM4 MSA with FEC",
	Ethernet100gPsm4:       "100G Ethernet: 100G PSM4 Parallel SMF",
	Ethernet100gAcc:        "100G Ethernet: 100G ACC or 25GAUI C2M ACC with worst BER of 5x10^(-5)",
	Ethernet100gCwdm4NoFec: "100G Ethernet: 100G CWDM4 MSA without FEC",
	Ethernet100gRsvd1:      "Reserved or unknown",
	Ethernet100gCr4:        "100G Ethernet: 100G Base-CR4 or 25G Base-CR CA-L",
	Ethernet25gCrCaS:       "25G Ethernet: 25G Base-CR CA-S",
	Ethernet25gCrCaN:       "25G Ethernet: 25G Base-CR CA-N",
	Ethernet40gEr4:         "40G Ethernet: 40G Base-ER4",
	Ethernet4X10Sr:         "4x10G Ethernet: 10G Base-SR",
	Ethernet40gPsm4:        "40G Ethernet: 40G PSM4 Parallel SMF",
	EthernetG959P1i1_2d1:   "Ethernet: G959.1 profile P1I1-2D1 (10709 MBd, 2km, 1310nm SM)",
	EthernetG959P1s1_2d2:   "Ethernet: G959.1 profile P1S1-2D2 (10709 MBd, 40km, 1550nm SM)",
	EthernetG959P1l1_2d2:   "Ethernet: G959.1 profile P1L1-2D2 (10709 MBd, 80km, 1550nm SM)",
	Ethernet10GtSfi:        "10G Ethernet: 10G Base-T with SFI electrical interface",
	Ethernet100gClr4:       "100G Ethernet: 100G CLR4",
	Ethernet100gAoc2:       "100G Ethernet: 100G AOC or 25GAUI C2M AOC with worst BER of 10^(-12)",
	Ethernet100gAcc2:       "100G Ethernet: 100G ACC or 25GAUI C2M ACC with worst BER of 10^(-12)",
}

var names = map[uint64]string{
	Ethernet10gLrm:    "10G Ethernet: 10G Base-LRM",
	Ethernet10gLr:     "10G Ethernet: 10G Base-LR",
	Ethernet10gSr:     "10G Ethernet: 10G Base-SR",
	Ethernet40gCr4:    "40G Ethernet: 40G Base-CR4",
	Ethernet40gSr4:    "40G Ethernet: 40G Base-SR4",
	Ethernet40gLr4:    "40G Ethernet: 40G Base-LR4",
	Ethernet40gActive: "40G Ethernet: 40G Active Cable (XLPPI)",
	Sonet40gOtn:       "40G OTN (OTU3B/OTU3C)",
	SonetOc48Lr:       "SONET: OC-48, long reach",
	SonetOc48Ir:       "SONET: OC-48, intermediate reach",
	SonetOc48Sr:       "SONET: OC-48, short reach",
	Sas6g:             "SAS 6.0G",
	Sas3g:             "SAS 3.0G",
	Gige1000BaseT:     "Ethernet: 1000BASE-T",
	Gige1000BaseCx:    "Ethernet: 1000BASE-CX",
	Gige1000BaseLx:    "Ethernet: 1000BASE-LX",
	Gige1000BaseSx:    "Ethernet: 1000BASE-SX",
	FcLenVeryLong:     "FC: very long distance (V)",
	FcLenShort:        "FC: short distance (S)",
	FcLenInt:          "FC: intermediate distance (I)",
	FcLenLong:         "FC: long distance (L)",
	FcLenMed:          "FC: medium distance (M)",
	FcTechLongLc:      "FC: Longwave laser (LC)",
	FcTechElecInter:   "FC: Electrical inter-enclosure (EL)",
	FcTechElecIntra:   "FC: Electrical intra-enclosure (EL)",
	FcTechShortWoOfc:  "FC: Shortwave laser w/o OFC (SN)",
	FcTechShortWOfc:   "FC: Shortwave laser with OFC (SL)",
	FcTechLongLl:      "FC: Longwave laser (LL)",
	FcTransMeidaTw:    "FC: Twin Axial Pair (TW)",
	FcTransMediaTp:    "FC: Twisted Pair (TP)",
	FcTransMediaMi:    "FC: Miniature Coax (MI)",
	FcTransMediaTv:    "FC: Video Coax (TV)",
	FcTransMediaM6:    "FC: Multimode, 62.5m (M6)",
	FcTransMediaM5:    "FC: Multimode, 50m (M5)",
	FcTransMediaOm3:   "FC: Multimode, 50um (OM3)",
	FcTransMediaSm:    "FC: Single Mode (SM)",
	FcSpeed1200Mbps:   "FC: 1200 Mb/s",
	FcSpeed800Mbps:    "FC: 800 Mb/s",
	FcSpeed1600Mbps:   "FC: 1600 Mb/s",
	FcSpeed400Mbps:    "FC: 400 Mb/s",
	FcSpeed200Mbps:    "FC: 200 Mb/s",
	FcSpeed100Mbps:    "FC: 100 Mb/s",
}

type uint64arr []uint64

func (a uint64arr) Len() int           { return len(a) }
func (a uint64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint64arr) Less(i, j int) bool { return a[i] < a[j] }

type Transceiver [8]byte

func (t Transceiver) List() []string {
	r := []string{}

	keys := uint64arr{}
	for k := range names {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	for _, k := range keys {
		if k&t.Uint64() != 0 {
			r = append(r, names[k])
		}
	}

	return r
}

func (t Transceiver) Uint64() uint64 {
	return *(*uint64)(unsafe.Pointer(&t[0]))
}

func (t Transceiver) String() string {
	return strings.Join(t.List(), "\n")
}

func (t Transceiver) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"values": t.List(),
		"hex":    hex.EncodeToString(t[:8]),
	}
	return json.Marshal(m)
}

func (t *Transceiver) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 8 {
		return fmt.Errorf("length is shorter then Transceiver type")
	}

	*t = Transceiver{}
	for i := 0; i < 8; i++ {
		t[i] = b[i]
	}
	return nil
}
