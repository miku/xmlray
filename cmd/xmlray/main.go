// xmlray is a little x-ray things for xml.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/miku/xmlray"
)

const Version = "0.0.1"

func main() {

	visitorName := flag.String("visitor", "default", "name of visitor to use")
	path := flag.String("path", "", "path to use for compact visitor")
	verbose := flag.Bool("verbose", false, "be verbose")
	version := flag.Bool("v", false, "show version and exit")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	var rdr io.Reader
	if flag.NArg() == 0 {
		rdr = os.Stdin
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		rdr = file
	}

	var visitor xmlray.Visitor

	switch *visitorName {
	case "d", "default":
		visitor = xmlray.VisitorFunc(func(s string) error {
			fmt.Println(s)
			return nil
		})
	case "c", "compact":
		visitor = xmlray.NewCompactVisitor(*path)
	case "s", "schema":
		visitor = xmlray.NewSchemaVisitor(*path, *verbose)
	default:
		log.Fatal("unknown visitor, use: default, compact, schema")
	}

	if err := xmlray.VisitNodes(rdr, visitor); err != nil {
		log.Fatal(err)
	}
}
