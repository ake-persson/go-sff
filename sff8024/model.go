package sff8024

import (
	"encoding/json"
)

type SFF8024 struct {
	Identifier Identifier `json:"identifier"` // 0 - Identifier
}

func (s *SFF8024) JSON() []byte {
	b, _ := json.Marshal(s)
	return b
}

func (s *SFF8024) JSONPretty() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return b
}
