package sff

import (
	"encoding/json"
	"errors"

	"github.com/mickep76/go-sff/sff8079"
	"github.com/mickep76/go-sff/sff8636"
)

type Type int

const (
	TypeUnknown Type = iota
	TypeSff8079
	TypeSff8636
)

type Module struct {
	Type Type `json:"type"`

	*sff8079.Sff8079
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

func (m *Module) MarshalJSON() ([]byte, error) {
	switch m.Type {
	case TypeSff8079:
		return json.Marshal(m.Sff8079)
	case TypeSff8636:
		return json.Marshal(m.Sff8636)
	}
	return nil, errors.New("unknown type")
}

/*
func (e *Encoding) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*e = Encoding(b[0])
	return nil
}
*/

func GetType(eeprom []byte) Type {
	switch len(eeprom) {
	case 256:
		return TypeSff8079
	case 640:
		return TypeSff8636
	}
	return TypeUnknown
}

func New(eeprom []byte) (*Module, error) {
	switch GetType(eeprom) {
	case TypeSff8079:
		m, err := sff8079.New(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8079, Sff8079: m}, nil
	case TypeSff8636:
		m, err := sff8636.New(eeprom)
		if err != nil {
			return nil, err
		}
		return &Module{Type: TypeSff8636, Sff8636: m}, nil
	}
	return nil, errors.New("unknown type")
}
