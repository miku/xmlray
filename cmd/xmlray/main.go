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

	"golang.org/x/net/html/charset"
)

const Version = "0.0.3"

// Visit lets a Visitor v visit all nodes in a XML doc wrapped in a reader.
func VisitReader(r io.Reader, v xmlray.NodeVisitor) error {
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
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
	return v.Visit(nil)
}

func main() {

	path := flag.String("path", "", "path to element of interest")
	vtype := flag.String("type", "path", "visitor type")
	version := flag.Bool("v", false, "show version and exit")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	visitors := map[string]xmlray.NodeVisitor{
		"string": xmlray.ChardataExtractor{},
		"debug":  xmlray.DebugVisitor{},
		"path":   &xmlray.PathVisitor{},
		"ns":     &xmlray.NamespaceLister{},
		"tag":    &xmlray.TagnameLister{},
		"group":  &xmlray.GroupVisitor{PathPrefix: *path},
		"tree":   xmlray.NewTreeVisitor(*path),
		"raw":    xmlray.NewRawVisitor(*path),
	}

	var reader io.Reader

	if flag.NArg() == 0 {
		reader = os.Stdin
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		reader = file
	}

	visitor, ok := visitors[*vtype]
	if !ok {
		log.Fatal("unknown visitor")
	}

	if err := VisitReader(bufio.NewReader(reader), visitor); err != nil {
		log.Fatal(err)
	}
}
