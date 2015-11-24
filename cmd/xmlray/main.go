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

	xmlray.VisitElements(rdr, func(s string, err error) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	})
}
