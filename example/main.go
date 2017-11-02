package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mickep76/go-sff"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	toJSON := flag.Bool("to-json", false, "Output as JSON")
	fromJSON := flag.Bool("from-json", false, "Input from JSON")
	flag.Parse()

	var b []byte
	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
	} else {
		log.Fatalf("stdin is hungry, feed me")
	}

	if *fromJSON {
		m := &sff.Module{}
		err := json.Unmarshal(b, m)
		if err != nil {
			log.Fatal(err)
		}

		if *toJSON {
			b, _ := json.MarshalIndent(m, "", "  ")
			fmt.Printf("%s\n", string(b))
		} else {
			fmt.Printf("%s\n", m.String())
		}
		return
	}

	eeprom, err := hex.DecodeString(strings.TrimRight(string(b), "\n"))
	if err != nil {
		log.Fatalf("decode hex: %v", err)
	}

	m, err := sff.Decode(eeprom)
	if err != nil {
		log.Fatal(err)
	}

	if *toJSON {
		b, _ := json.MarshalIndent(m, "", "  ")
		fmt.Printf("%s\n", string(b))
	} else {
		fmt.Printf("%s\n", m.String())
	}
}
