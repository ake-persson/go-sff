package sff

const (
	TypeUnknown Type = iota
	TypeSff8079
	TypeSff8636
)

type Type int

func GetType(eeprom []byte) Type {
	switch len(eeprom) {
	case 256:
		return TypeSff8079
	case 640:
		return TypeSff8636
	default:
		return TypeUnknown
	}
}
