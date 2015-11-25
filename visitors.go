package xmlray

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

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
	Path    string
	Verbose bool
	// seen keeps track of all observed paths, elements, attributes and chardata nodes
	seen map[string]bool
	// repeatable keeps track, which elements are observed multiple time within a unit
	repeatable map[string]bool
	m          map[string]int
}

// NewCompactVisitor returns a new compact visitor, given a path element, that
// is taken as the root element.
func NewSchemaVisitor(s string, verbose bool) *SchemaVisitor {
	return &SchemaVisitor{Path: s,
		Verbose:    verbose,
		m:          make(map[string]int),
		seen:       make(map[string]bool),
		repeatable: make(map[string]bool)}
}

// Visit visits nodes and keeps track of how often a particular type has been
// seen.
func (v SchemaVisitor) Visit(s string) error {
	if v.Verbose {
		log.Printf("%+v -- %s", v, s)
	}
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
	if v.Verbose {
		log.Println("flushing visitor")
	}

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
