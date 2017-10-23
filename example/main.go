package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mickep76/go-sff"
	"github.com/mickep76/go-sff/sff8079"
	"github.com/mickep76/go-sff/sff8636"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	printAsJSON := flag.Bool("json", false, "Print output as JSON")
	flag.Parse()

	var b []byte
	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
	} else {
		log.Fatalf("stdin is hungry, feed me")
	}

	eeprom, err := hex.DecodeString(strings.TrimRight(string(b), "\n"))
	if err != nil {
		log.Fatalf("decode hex: %v", err)
	}

	switch sff.GetType(eeprom) {
	case sff.TypeSff8079:
		m, _ := sff8079.New(eeprom)

		if *printAsJSON {
			fmt.Printf("%s\n", m.JSONPretty())
		} else {
			fmt.Printf("%s\n", m)
		}
	case sff.TypeSff8636:
		m, _ := sff8636.New(eeprom)

		if *printAsJSON {
			fmt.Printf("%s\n", m.JSONPretty())
		} else {
			fmt.Printf("%s\n", m)
		}
	default:
		log.Fatal("unknown eeprom type")
	}
}
