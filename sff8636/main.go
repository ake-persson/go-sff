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
	"golang.org/x/crypto/ssh/terminal"
)

type ByteString2 [2]byte
type ByteString16 [16]byte
type VendorOUI [3]byte
type DateCode [8]byte

func (b ByteString2) String() string {
	return strings.TrimSpace(string(b[0:2]))
}

func (b ByteString2) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b ByteString16) String() string {
	return strings.TrimSpace(string(b[0:16]))
}

func (b ByteString16) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (v VendorOUI) String() string {
	return fmt.Sprintf("%x:%x:%x", v[0], v[1], v[2])
}

func (v VendorOUI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (d DateCode) String() string {
	return fmt.Sprintf("20%s-%s-%s", string(d[0:2]), string(d[2:4]), string(d[4:6]))
}

func (d DateCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
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
		s8024 := (*sff8024.SFF8024)(unsafe.Pointer(&eeprom[0]))
		s8636 := (*SFF8636)(unsafe.Pointer(&eeprom[129]))
		fmt.Printf("%s\n", s8024.JSONPretty())
		fmt.Printf("%s\n", s8636.JSONPretty())
	}
}
