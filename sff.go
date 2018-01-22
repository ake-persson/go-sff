package sff

import (
	"encoding/json"
	"errors"

	"github.com/mickep76/go-sff/sff8079"
	"github.com/mickep76/go-sff/sff8636"
)

// Type of eeprom module.
type Type string

const (
	TypeUnknown = Type("Unknown")
	TypeSff8079 = Type("SFF-8079")
	TypeSff8636 = Type("SFF-8636")
)

var ErrUnknownType = errors.New("unknown type")

type Module struct {
	Type             Type `json:"type"`
	*sff8079.Sff8079 `json:"-"`
	*sff8636.Sff8636 `json:"-"`
}

type module Module

type moduleSff8079 struct {
	Type Type `json:"type"`
	*sff8079.Sff8079
}

type moduleSff8636 struct {
	Type Type `json:"type"`
	*sff8636.Sff8636
}

func (m *Module) String() string {
	switch m.Type {
	case TypeSff8079:
		return m.Sff8079.String()
	case TypeSff8636:
		return m.Sff8636.String()
	}
	return ""
}

func (m *Module) StringCol() string {
	switch m.Type {
	case TypeSff8079:
		return m.Sff8079.StringCol()
	case TypeSff8636:
		return m.Sff8636.StringCol()
	}
	return ""
}

func (m *Module) MarshalJSON() ([]byte, error) {
	switch m.Type {
	case TypeSff8079:
		return json.Marshal(moduleSff8079{Type: m.Type, Sff8079: m.Sff8079})
	case TypeSff8636:
		return json.Marshal(moduleSff8636{Type: m.Type, Sff8636: m.Sff8636})
	}
	return nil, ErrUnknownType
}

func (m *Module) UnmarshalJSON(in []byte) error {
	mod := &module{}
	err := json.Unmarshal(in, mod)
	if err != nil {
		return err
	}
	m.Type = mod.Type

	switch mod.Type {
	case TypeSff8079:
		s := &sff8079.Sff8079{}
		if err := json.Unmarshal(in, s); err != nil {
			return err
		}
		m.Sff8079 = s
		return nil
	case TypeSff8636:
		s := &sff8636.Sff8636{}
		if err := json.Unmarshal(in, s); err != nil {
			return err
		}
		m.Sff8636 = s
		return nil
	}
	return ErrUnknownType
}

func GetType(eeprom []byte) Type {
	if (eeprom[0] == 2 || eeprom[0] == 3) && eeprom[1] == 4 {
		return TypeSff8079
	}

	if eeprom[128] == 12 || eeprom[128] == 13 || eeprom[128] == 17 {
		return TypeSff8636
	}

	return TypeUnknown
}

func Decode(eeprom []byte) (*Module, error) {
	switch GetType(eeprom) {
	case TypeSff8079:
		m, err := sff8079.Decode(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8079, Sff8079: m}, nil
	case TypeSff8636:
		m, err := sff8636.Decode(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8636, Sff8636: m}, nil
	}
	return nil, ErrUnknownType
}
