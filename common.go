package xmlray

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Visitor interface {
	Visit(string) error
}

// VisitorFunc will be called with the current path string and the XML element.
type VisitorFunc func(string) error

func (f VisitorFunc) Visit(s string) error {
	return f(s)
}

type Stack []string

func (s Stack) String() string {
	return "/" + strings.Join(s, "/")
}

func VisitElements(r io.Reader, v Visitor) error {
	dec := xml.NewDecoder(r)
	var stack Stack
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			if err := v.Visit(stack.String()); err != nil {
				return err
			}
			for _, attr := range tok.Attr {
				if err := v.Visit(stack.String() + "/@" + attr.Name.Local); err != nil {
					return err
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			cleaned := strings.TrimSpace(string(tok))
			if cleaned != "" {
				if err := v.Visit(stack.String() + "/#"); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CompactVisitor struct {
	Path string
	m    map[string]int
}

func NewCompactVisitor(s string) *CompactVisitor {
	return &CompactVisitor{Path: s, m: make(map[string]int)}
}

func (v CompactVisitor) Visit(s string) error {
	if s == v.Path {
		if v.m != nil {
			for k, v := range v.m {
				fmt.Println(k, v)
			}
			fmt.Println("--")
		}
		for k := range v.m {
			delete(v.m, k)
		}
	}
	v.m[s]++
	return nil
}
