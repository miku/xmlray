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
	Repeatable      bool            `json:"repeatable"`
	HasChardata     bool            `json:"hasChardata"`
	ChardataSamples []string        `json:"samples"`
	AttributeNames  map[string]bool `json:"attributeNames"`
	// add more attributes as needed
}

// GroupVisitor groups elements starting at PathPrefix.
type GroupVisitor struct {
	PathPrefix string
	stack      []string
	nodeNames  []string
	recording  bool
	// tagMap contains the global information about the
	// tagMap map[string]tagInfo
	localMap map[string]tagInfo
}

func (v *GroupVisitor) path() string {
	return "/" + strings.Join(v.stack, "/")
}

// updateMapping updates the globale entry map.
func (v *GroupVisitor) updateMapping() {
	b, err := json.Marshal(v.localMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	v.localMap = nil

	// for _, name := range v.nodeNames {
	// 	log.Println(name)
	// }
	// log.Println("----")
}

func (v *GroupVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.CharData:
		if v.recording {
			last := v.nodeNames[len(v.nodeNames)-1]

			ti, ok := v.localMap[last]
			if !ok {
				panic("invalid localMap")
			}

			s := strings.TrimSpace(string(node))
			if s == "" {
				return nil
			}

			ti.HasChardata = true
			ti.ChardataSamples = append(ti.ChardataSamples, s)
			v.localMap[v.path()] = ti

			// log.Println("last", last)

			// if strings.Contains(last, "@") {
			// 	return nil
			// }
			// if last != "char" && strings.TrimSpace(string(node)) != "" {
			// 	v.nodeNames = append(v.nodeNames, last+"/#")
			// }
		}
	case xml.StartElement:
		v.stack = append(v.stack, node.Name.Local)
		if strings.HasPrefix(v.path(), v.PathPrefix) {
			v.recording = true

			if v.localMap == nil {
				v.localMap = make(map[string]tagInfo)
			}

			v.nodeNames = append(v.nodeNames, v.path())
			// for _, attr := range node.Attr {
			// 	v.nodeNames = append(v.nodeNames, v.path()+"/@"+attr.Name.Local)
			// }

			// the current nodes attribute names
			attrNames := make(map[string]bool)
			for _, attr := range node.Attr {
				attrNames[attr.Name.Local] = true
			}

			// lookup tag info if there is any
			ti, ok := v.localMap[v.path()]

			// if there is none, create a new one, otherwise, this element is repeatable
			if !ok {
				v.localMap[v.path()] = tagInfo{
					Repeatable:      false,
					AttributeNames:  attrNames,
					ChardataSamples: make([]string, 0),
				}
			} else {
				ti.Repeatable = true
				for name := range attrNames {
					ti.AttributeNames[name] = true
				}
				v.localMap[v.path()] = ti
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

type TreeVisitor struct {
	PathPrefix string
	stack      []string
	recording  bool
	// store info per node
	nodes map[string]interface{}
	// childmap
	childmap NodeMap
}

type NodeMap map[string][]string

func (m NodeMap) Add(parent, child string) {
	m[parent] = append(m[parent], child)
}

func NewTreeVisitor(path string) *TreeVisitor {
	return &TreeVisitor{
		PathPrefix: path,
		nodes:      make(map[string]interface{}),
		childmap:   make(NodeMap),
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

func findNamespaces(attr []xml.Attr) map[string]string {
	var nss = make(map[string]string)
	for _, at := range attr {
		if at.Name.Space == "xmlns" {
			nss[at.Name.Local] = at.Value
		}
	}
	return nss
}

func findAttributes(attr []xml.Attr) map[string]string {
	var attributes = make(map[string]string)
	for _, at := range attr {
		if at.Name.Space != "xmlns" {
			attributes[at.Name.Local] = at.Value
		}
	}
	return attributes
}

func (v *TreeVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.StartElement:
		if v.path() == v.PathPrefix {
			v.recording = true
		}
		v.stack = append(v.stack, node.Name.Local)
		if v.recording {
			v.childmap.Add(v.parent(), v.path())
			v.nodes[v.path()] = map[string]interface{}{
				"nss":   findNamespaces(node.Attr),
				"attr":  findAttributes(node.Attr),
				"name":  node.Name.Local,
				"space": node.Name.Space,
			}
		}
	case xml.EndElement:
		if v.path() == v.PathPrefix {
			v.recording = false
			b, err := json.Marshal(v.childmap)
			if err != nil {
				return err
			}
			fmt.Println(string(b))

			b, err = json.Marshal(v.nodes)
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
