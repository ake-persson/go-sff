package main

import (
	"encoding/hex"
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
		s8024 := (*sff8024.SFF8024)(unsafe.Pointer(&eeprom[0]))
		s8636 := (*sff8636.SFF8636)(unsafe.Pointer(&eeprom[129]))
		fmt.Printf("%s\n", s8024.JSONPretty())
		fmt.Printf("%s\n", s8636.JSONPretty())
	}
}
