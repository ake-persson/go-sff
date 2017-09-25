package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/mickep76/go-sff/sff8079"
	"github.com/mickep76/go-sff/sff8636"
	"golang.org/x/crypto/ssh/terminal"
)

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

	switch len(eeprom) {
	case 256:
		m := (*sff8079.SFF8079)(unsafe.Pointer(&eeprom[0]))
		fmt.Printf("%s\n", m.JSONPretty())
	case 640:
		m := (*sff8636.SFF8636)(unsafe.Pointer(&eeprom[128]))
		fmt.Printf("%s\n", m.JSONPretty())
	default:
		log.Fatal("unknown eeprom size: %d", len(eeprom))
	}
}
