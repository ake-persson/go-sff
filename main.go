package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/mickep76/go-sff/sff8024"
	"github.com/mickep76/go-sff/sff8636"
	"golang.org/x/crypto/ssh/terminal"
)

type Module struct {
	*sff8024.SFF8024
	*sff8636.SFF8636
}

func (m *Module) JSON() []byte {
	b, _ := json.Marshal(m)
	return b
}

func (m *Module) JSONPretty() []byte {
	b, _ := json.MarshalIndent(m, "", "  ")
	return b
}

func main() {
	var b []byte
	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
	} else {
		log.Fatalf("stdin is hungry, feed me")
	}

	// Decode hex
	eeprom, err := hex.DecodeString(strings.TrimRight(string(b), "\n"))
	if err != nil {
		log.Fatalf("decode hex: %v", err)
	}
	fmt.Printf("Eeprom Size: %d\n", len(eeprom))

	if len(eeprom) == 512 || len(eeprom) == 640 {
		m := Module{
			SFF8024: (*sff8024.SFF8024)(unsafe.Pointer(&eeprom[0])),
			SFF8636: (*sff8636.SFF8636)(unsafe.Pointer(&eeprom[129])),
		}
		fmt.Printf("%s\n", m.JSONPretty())
	}
}
