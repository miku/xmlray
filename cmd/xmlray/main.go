package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/miku/xmlray"
)

func main() {

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

	// visitor := xmlray.VisitorFunc(func(s string) error {
	// 	fmt.Println(s)
	// 	return nil
	// })

	visitor := xmlray.NewCompactVisitor("/article")
	err := xmlray.VisitElements(rdr, visitor)

	if err != nil {
		log.Fatal(err)
	}
}
