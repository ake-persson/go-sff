package sff8079

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

const (
	Ether10gBaseEr      = (1 << 7)
	Ether10gBaseLrm     = (1 << 6)
	Ether10gBaeLr       = (1 << 5)
	Ether10gBaseSr      = (1 << 4)
	Infini1xSx          = (1 << 3)
	Infini1xLx          = (1 << 2)
	Infini1xCopprActv   = (1 << 1)
	Infini1xCopprPasv   = (1 << 0)
	EsconMmf            = (1 << (7 + 8))
	EsconSmf            = (1 << (6 + 8))
	SonetOc192Short     = (1 << (5 + 8))
	SonetReachBit1      = (1 << (4 + 8))
	SonetReachBit2      = (1 << (3 + 8))
	SonetOc48Long       = (1 << (2 + 8))
	SonetOc48Inter      = (1 << (1 + 8))
	SonetOc48Short      = (1 << (0 + 8))
	SonetOc12Long       = (1 << (6 + 16))
	SonetOc12Inter      = (1 << (5 + 16))
	SonetOc12Short      = (1 << (4 + 16))
	SonetOc3Long        = (1 << (2 + 16))
	SonetOc3Inter       = (1 << (1 + 16))
	SonetOc3Short       = (1 << (0 + 16))
	EtherBasePx         = (1 << (7 + 24))
	EtherBaseBx10       = (1 << (6 + 24))
	Ether100BaseFx      = (1 << (5 + 24))
	Ether100BaseLx      = (1 << (4 + 24))
	Ether1000BaseT      = (1 << (3 + 24))
	Ether1000BaseCx     = (1 << (2 + 24))
	Ether1000BaseLx     = (1 << (1 + 24))
	Ether1000BaseSx     = (1 << (0 + 24))
	FcVeryLongDist      = (1 << (7 + 32))
	FcShortDist         = (1 << (6 + 32))
	FcIntermDist        = (1 << (5 + 32))
	FcLongDist          = (1 << (4 + 32))
	FcMedDist           = (1 << (3 + 32))
	FcShortwaveLaser    = (1 << (2 + 32))
	FcLongwaveLaser     = (1 << (1 + 32))
	FcElectrInterEncl   = (1 << (0 + 32))
	FcElecIntraEncl     = (1 << (7 + 40))
	FcShortwaveWoOfc    = (1 << (6 + 40))
	FcShortwaveLaserOfc = (1 << (5 + 40))
	FcLongwaveLasterLl  = (1 << (4 + 40))
	ActiveCable         = (1 << (3 + 40))
	PassiveCable        = (1 << (2 + 40))
	FcCopperBaseT       = (1 << (1 + 40))
	FcTwinAxialPair     = (1 << (7 + 48))
	FcTwistedPair       = (1 << (6 + 48))
	FcMiniCoaxMi        = (1 << (5 + 48))
	FcVideoCoaxTv       = (1 << (4 + 48))
	FcMultimodeM6       = (1 << (3 + 48))
	FcMultimodeM5       = (1 << (2 + 48))
	FcSingleMode        = (1 << (0 + 48))
	Fc1200MbPerSec      = (1 << (7 + 56))
	Fc800MbPerSec       = (1 << (6 + 56))
	Fc400MbPerSec       = (1 << (4 + 56))
	Fc200MbPerSec       = (1 << (2 + 56))
	Fc100MbPerSec       = (1 << (0 + 56))
)

var transceiverNames = map[uint64]string{
	Ether10gBaseEr:      "10G Ethernet: 10G Base-ER [SFF-8472 rev10.4 only]",
	Ether10gBaseLrm:     "10G Ethernet: 10G Base-LRM",
	Ether10gBaeLr:       "10G Ethernet: 10G Base-LR",
	Ether10gBaseSr:      "10G Ethernet: 10G Base-SR",
	Infini1xSx:          "Infiniband: 1X SX",
	Infini1xLx:          "Infiniband: 1X LX",
	Infini1xCopprActv:   "Infiniband: 1X Copper Active",
	Infini1xCopprPasv:   "Infiniband: 1X Copper Passive",
	EsconMmf:            "ESCON: ESCON MMF, 1310nm LED",
	EsconSmf:            "ESCON: ESCON SMF, 1310nm Laser",
	SonetOc192Short:     "SONET: OC-192, short reach",
	SonetReachBit1:      "SONET: SONET reach specifier bit 1",
	SonetReachBit2:      "SONET: SONET reach specifier bit 2",
	SonetOc48Long:       "SONET: OC-48, long reach",
	SonetOc48Inter:      "SONET: OC-48, intermediate reach",
	SonetOc48Short:      "SONET: OC-48, short reach",
	SonetOc12Long:       "SONET: OC-12, single mode, long reach",
	SonetOc12Inter:      "SONET: OC-12, single mode, inter. reach",
	SonetOc12Short:      "SONET: OC-12, short reach",
	SonetOc3Long:        "SONET: OC-3, single mode, long reach",
	SonetOc3Inter:       "SONET: OC-3, single mode, inter. reach",
	SonetOc3Short:       "SONET: OC-3, short reach",
	EtherBasePx:         "Ethernet: BASE-PX",
	EtherBaseBx10:       "Ethernet: BASE-BX10",
	Ether100BaseFx:      "Ethernet: 100BASE-FX",
	Ether100BaseLx:      "Ethernet: 100BASE-LX/LX10",
	Ether1000BaseT:      "Ethernet: 1000BASE-T",
	Ether1000BaseCx:     "Ethernet: 1000BASE-CX",
	Ether1000BaseLx:     "Ethernet: 1000BASE-LX",
	Ether1000BaseSx:     "Ethernet: 1000BASE-SX",
	FcVeryLongDist:      "FC: very long distance (V)",
	FcShortDist:         "FC: short distance (S)",
	FcIntermDist:        "FC: intermediate distance (I)",
	FcLongDist:          "FC: long distance (L)",
	FcMedDist:           "FC: medium distance (M)",
	FcShortwaveLaser:    "FC: Shortwave laser, linear Rx (SA)",
	FcLongwaveLaser:     "FC: Longwave laser (LC)",
	FcElectrInterEncl:   "FC: Electrical inter-enclosure (EL)",
	FcElecIntraEncl:     "FC: Electrical intra-enclosure (EL)",
	FcShortwaveWoOfc:    "FC: Shortwave laser w/o OFC (SN)",
	FcShortwaveLaserOfc: "FC: Shortwave laser with OFC (SL)",
	FcLongwaveLasterLl:  "FC: Longwave laser (LL)",
	ActiveCable:         "Active Cable",
	PassiveCable:        "Passive Cable",
	FcCopperBaseT:       "FC: Copper FC-BaseT",
	FcTwinAxialPair:     "FC: Twin Axial Pair (TW)",
	FcTwistedPair:       "FC: Twisted Pair (TP)",
	FcMiniCoaxMi:        "FC: Miniature Coax (MI)",
	FcVideoCoaxTv:       "FC: Video Coax (TV)",
	FcMultimodeM6:       "FC: Multimode, 62.5um (M6)",
	FcMultimodeM5:       "FC: Multimode, 50um (M5)",
	FcSingleMode:        "FC: Single Mode (SM)",
	Fc1200MbPerSec:      "FC: 1200 MBytes/sec",
	Fc800MbPerSec:       "FC: 800 MBytes/sec",
	Fc400MbPerSec:       "FC: 400 MBytes/sec",
	Fc200MbPerSec:       "FC: 200 MBytes/sec",
	Fc100MbPerSec:       "FC: 100 MBytes/sec",
}

type uint64arr []uint64

func (a uint64arr) Len() int           { return len(a) }
func (a uint64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint64arr) Less(i, j int) bool { return a[i] < a[j] }

type Transceiver [8]byte

func (t Transceiver) List() []string {
	r := []string{}

	keys := uint64arr{}
	for k := range transceiverNames {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	for _, k := range keys {
		if k&t.Uint64() != 0 {
			r = append(r, transceiverNames[k])
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
