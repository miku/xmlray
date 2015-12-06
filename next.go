package xmlray

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

const NoNameSpace = "none"

// NodeVisitor gets passed one of the XML nodes.
type NodeVisitor interface {
	Visit(interface{}) error
	Flush() error
}

// DebugVisitor displays node information.
type DebugVisitor struct{}

// Visits prints debug info on nodes.
func (v DebugVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.StartElement:
		log.Printf("%T %s:%s", node, node.Name.Space, node.Name.Local)
	case xml.EndElement:
		log.Printf("%T %s:%s", node, node.Name.Space, node.Name.Local)
	case xml.CharData:
		s := string(node)
		l := len(s)
		stripped := strings.TrimSpace(s)
		log.Printf("%T s='%s', %d/%d", node, stripped, l, len(stripped))
	}
	return nil
}

// Flush does nothing here.
func (v DebugVisitor) Flush() error {
	return nil
}

// ChardataExtractor extracts only strings.
type ChardataExtractor struct{}

// Visit extracts text, if node is CharData.
func (v ChardataExtractor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.CharData:
		s := strings.TrimSpace(string(node))
		if len(s) > 0 {
			fmt.Println(s)
		}
	}
	return nil
}

// Flush does nothing here.
func (v ChardataExtractor) Flush() error {
	return nil
}

// NamespaceLister collects namespace usage.
type NamespaceLister struct {
	ns map[string]int
}

// Visit collects the namespace from start elements.
func (v *NamespaceLister) Visit(node interface{}) error {
	if v.ns == nil {
		v.ns = make(map[string]int)
	}
	switch node := node.(type) {
	case xml.StartElement:
		switch node.Name.Space {
		case "":
			v.ns[NoNameSpace]++
		default:
			v.ns[node.Name.Space]++
		}
	}
	return nil
}

// Flush dumps result to stdout.
func (v *NamespaceLister) Flush() error {
	b, err := json.Marshal(v.ns)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

type TagnameLister struct {
	tags map[string]map[string]int
}

// Visit collects the namespace from start elements.
func (v *TagnameLister) Visit(node interface{}) error {
	if v.tags == nil {
		v.tags = make(map[string]map[string]int)
	}
	switch node := node.(type) {
	case xml.StartElement:
		switch node.Name.Space {
		case "":
			if v.tags[NoNameSpace] == nil {
				v.tags[NoNameSpace] = make(map[string]int)
			}
			v.tags[NoNameSpace][node.Name.Local]++
		default:
			if v.tags[node.Name.Space] == nil {
				v.tags[node.Name.Space] = make(map[string]int)
			}
			v.tags[node.Name.Space][node.Name.Local]++
		}
	}
	return nil
}

// Flush dumps result to stdout.
func (v *TagnameLister) Flush() error {
	b, err := json.Marshal(v.tags)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
