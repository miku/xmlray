package xmlray

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

const NoNameSpaceKey = "none"

// NodeVisitor gets passed one of the XML nodes. The sentinel value is nil.
type NodeVisitor interface {
	Visit(interface{}) error
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
			v.ns[NoNameSpaceKey]++
		default:
			v.ns[node.Name.Space]++
		}
	case nil:
		return v.flush()
	}
	return nil
}

// flush dumps result to stdout.
func (v *NamespaceLister) flush() error {
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
			if v.tags[NoNameSpaceKey] == nil {
				v.tags[NoNameSpaceKey] = make(map[string]int)
			}
			v.tags[NoNameSpaceKey][node.Name.Local]++
		default:
			if v.tags[node.Name.Space] == nil {
				v.tags[node.Name.Space] = make(map[string]int)
			}
			v.tags[node.Name.Space][node.Name.Local]++
		}
	case nil:
		return v.flush()
	}
	return nil
}

// flush dumps result to stdout.
func (v *TagnameLister) flush() error {
	b, err := json.Marshal(v.tags)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

// PathVisitor lists all path (from start elements).
type PathVisitor struct {
	stack []string
}

func (v *PathVisitor) path() string {
	return "/" + strings.Join(v.stack, "/")
}

func (v *PathVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.StartElement:
		v.stack = append(v.stack, node.Name.Local)
		if len(v.stack) > 0 {
			fmt.Println(v.path())
		}
	case xml.EndElement:
		v.stack = v.stack[:len(v.stack)-1]
	}
	return nil
}

type tagInfo struct {
	Repeatable     bool
	HasChardata    bool
	AttributeNames []string
}

// GroupVisitor groups elements starting at PathPrefix.
type GroupVisitor struct {
	PathPrefix string
	stack      []string
	nodeNames  []string
	recording  bool

	tagMap map[string]tagInfo
}

func (v *GroupVisitor) path() string {
	return "/" + strings.Join(v.stack, "/")
}

// updateMapping updates the globale entry map.
func (v *GroupVisitor) updateMapping() {
	for _, name := range v.nodeNames {
		log.Println(name)
	}
	log.Println("----")
}

func (v *GroupVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.CharData:
		if v.recording {
			last := v.nodeNames[len(v.nodeNames)-1]
			if strings.Contains(last, "@") {
				return nil
			}
			if last != "char" && strings.TrimSpace(string(node)) != "" {
				v.nodeNames = append(v.nodeNames, last+"/#")
			}
		}
	case xml.StartElement:
		v.stack = append(v.stack, node.Name.Local)
		if strings.HasPrefix(v.path(), v.PathPrefix) {
			v.recording = true
			v.nodeNames = append(v.nodeNames, v.path())
			for _, attr := range node.Attr {
				v.nodeNames = append(v.nodeNames, v.path()+"/@"+attr.Name.Local)
			}
		}
	case xml.EndElement:
		if v.path() == v.PathPrefix {
			v.updateMapping()
			v.nodeNames = v.nodeNames[:0]
			v.recording = false
		}
		v.stack = v.stack[:len(v.stack)-1]
		return nil
	default:
		return nil
	}
	return nil
}

type info struct{}

type TreeVisitor struct {
	PathPrefix string
	stack      []string
	recording  bool
	// store info per node
	nodes map[string]info
	// childmap
	childmap map[string][]string
}

func NewTreeVisitor(path string) *TreeVisitor {
	return &TreeVisitor{
		PathPrefix: path,
		nodes:      make(map[string]info),
		childmap:   make(map[string][]string),
	}
}

func (v *TreeVisitor) path() string {
	return "/" + strings.Join(v.stack, "/")
}

func (v *TreeVisitor) parent() string {
	if len(v.stack) < 2 {
		return ""
	}
	return "/" + strings.Join(v.stack[:len(v.stack)-1], "/")
}

func (v *TreeVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.StartElement:
		if v.path() == v.PathPrefix {
			v.recording = true
		}
		v.stack = append(v.stack, node.Name.Local)
		if v.recording {
			v.childmap[v.parent()] = append(v.childmap[v.parent()], v.path())
		}
	case xml.EndElement:
		if v.path() == v.PathPrefix {
			v.recording = false
			b, err := json.Marshal(v.childmap)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			v.childmap = make(map[string][]string)
		}
		v.stack = v.stack[:len(v.stack)-1]
		return nil

	}
	return nil
}
