// xmlray provides xml node visiting functions to be used with the command line tool xmlray.
package xmlray

import (
	"encoding/xml"
	"io"
	"strings"
)

// Visitor is a simpQle visitor, that works with strings.
type Visitor interface {
	// Visit gets passed a XML path string.
	Visit(string) error
	// Flush is necessary, only if the visitor accumulates some state.
	Flush() error
}

// VisitorFunc can turn functions into Visitors.
type VisitorFunc func(string) error

// Visit visits nodes.
func (f VisitorFunc) Visit(s string) error {
	return f(s)
}

// Flush flushes state.
func (f VisitorFunc) Flush() error {
	return nil
}

// Stack is a simple stack.
type Stack []string

// String returns a path representation of the elements in the stack.
func (s Stack) String() string {
	return "/" + strings.Join(s, "/")
}

// VisitNodes lets a Visitor v visit most nodes in a XML doc wrapped in a
// reader.
func VisitNodes(r io.Reader, v Visitor) error {
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
	return v.Flush()
}
