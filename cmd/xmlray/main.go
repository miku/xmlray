package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/miku/xmlray"
)

func main() {

	visitorName := flag.String("visitor", "default", "name of visitor to use")
	path := flag.String("path", "", "path to use for compact visitor")

	flag.Parse()

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
	case "default":
		visitor = xmlray.VisitorFunc(func(s string) error {
			fmt.Println(s)
			return nil
		})
	case "c", "compact":
		visitor = xmlray.NewCompactVisitor(*path)
	default:
		log.Fatal("unknown visitor, use: default or compact")
	}

	err := xmlray.VisitElements(rdr, visitor)
	if err != nil {
		log.Fatal(err)
	}
}
