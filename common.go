package xmlray

import (
	"encoding/xml"
	"fmt"
	"io"
	"sort"
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

// CompactVisitor keeps track how many elements of a particular type have been
// observed.
type CompactVisitor struct {
	Path string
	m    map[string]int
	sep  string
}

// NewCompactVisitor returns a new compact visitor, given a path element, that
// is taken as the root element.
func NewCompactVisitor(s string) *CompactVisitor {
	return &CompactVisitor{Path: s, m: make(map[string]int), sep: "----"}
}

// Visit visits nodes and keeps track of how often a particular type has been
// seen.
func (v CompactVisitor) Visit(s string) error {
	if len(s) < len(v.Path) {
		return nil
	}
	if s == v.Path {
		for k, v := range v.m {
			fmt.Println(k, v)
		}
		fmt.Println(v.sep)
		for k := range v.m {
			delete(v.m, k)
		}
	}
	v.m[s]++
	return nil
}

// Flush prints out the remaining. Necessary, because only StartElement events
// are observed. TODO(miku): observe all events?
func (v CompactVisitor) Flush() error {
	for k, v := range v.m {
		fmt.Println(k, v)
	}
	return nil
}

// SchemaVisitor helps infering a simple schema.
type SchemaVisitor struct {
	Path string
	// seen keeps track of all observed paths, elements, attributes and chardata nodes
	seen map[string]bool
	// repeatable keeps track, which elements are observed multiple time within a unit
	repeatable map[string]bool
	m          map[string]int
}

// NewCompactVisitor returns a new compact visitor, given a path element, that
// is taken as the root element.
func NewSchemaVisitor(s string) *SchemaVisitor {
	return &SchemaVisitor{Path: s,
		m:          make(map[string]int),
		seen:       make(map[string]bool),
		repeatable: make(map[string]bool)}
}

// Visit visits nodes and keeps track of how often a particular type has been
// seen.
func (v SchemaVisitor) Visit(s string) error {
	if !strings.HasPrefix(s, v.Path) {
		return nil
	}
	if s == v.Path {
		for k, count := range v.m {
			v.seen[k] = true
			if count > 1 {
				v.repeatable[k] = true
			}
		}
		for k := range v.m {
			delete(v.m, k)
		}
	}
	v.m[s]++
	return nil
}

// Flush prints out the remaining. Necessary, because only StartElement events
// are observed. TODO(miku): observe all events?
func (v SchemaVisitor) Flush() error {
	var lines []string
	for k, _ := range v.seen {
		_, found := v.repeatable[k]
		if found {
			lines = append(lines, fmt.Sprintf("[]"+k))
		} else {
			lines = append(lines, fmt.Sprintf(k))
		}
	}
	sort.Strings(lines)
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
