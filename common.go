package xmlray

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

// VisitorFunc will be called with the current path string and the XML element.
type VisitorFunc func(string, error)

type Stack []string

func (s Stack) String() string {
	return "/" + strings.Join(s, "/")
}

func VisitElements(r io.Reader, visit VisitorFunc) {
	dec := xml.NewDecoder(r)
	var stack Stack
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			visit(stack.String(), err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			visit(stack.String(), nil)
			for _, attr := range tok.Attr {
				visit(stack.String()+"/@"+attr.Name.Local, nil)
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			cleaned := strings.TrimSpace(string(tok))
			if cleaned != "" {
				visit(stack.String()+"/#", nil)
			}
		}
	}
}

func PrintVisitor(s string, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
