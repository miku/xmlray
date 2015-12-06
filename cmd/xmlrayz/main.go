// xmlray is a little x-ray things for xml.
package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/miku/xmlray"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

const Version = "0.0.2"

// Visit lets a Visitor v visit most nodes in a XML doc wrapped in a
// reader.
func VisitReader(r io.Reader, v xmlray.NodeVisitor) error {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReader
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := v.Visit(tok); err != nil {
			return err
		}
	}
	return v.Flush()
}

func main() {

	version := flag.Bool("v", false, "show version and exit")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	var r io.Reader

	if flag.NArg() == 0 {
		r = os.Stdin
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		r = file
	}

	br := bufio.NewReader(r)
	if err := VisitReader(br, &xmlray.NamespaceLister{}); err != nil {
		log.Fatal(err)
	}
}
