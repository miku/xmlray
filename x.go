package xmlray

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

// ChildMap stores the the child elements per parent. No namespace support.
type ChildMap struct {
	nodes map[string][]string
	attrs map[string][]string
}

func NewChildMap() *ChildMap {
	return &ChildMap{
		nodes: make(map[string][]string),
		attrs: make(map[string][]string),
	}
}

func (cm *ChildMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"nodes": cm.nodes,
		"attrs": cm.attrs,
	})
}

// AddNode adds a childnode.
func (cm *ChildMap) AddNode(parent, child string) {
	cm.nodes[parent] = append(cm.nodes[parent], child)
}

// AddAttr adds an attribute.
func (cm *ChildMap) AddAttr(node, name string) {
	cm.attrs[node] = append(cm.attrs[node], name)
}

// StringStack is a simple stack.
type StringStack struct {
	fifo []string
}

// Push adds an element to the stack.
func (s *StringStack) Push(v string) {
	s.fifo = append(s.fifo, v)
}

// Pop removes an element, panics on empty stack.
func (s *StringStack) Pop() string {
	item := s.fifo[len(s.fifo)-1]
	s.fifo = s.fifo[:len(s.fifo)-1]
	return item
}

func (s *StringStack) Path() string {
	return fmt.Sprintf("/%s", strings.Join(s.fifo, "/"))
}

func (s *StringStack) Name() string {
	if len(s.fifo) == 0 {
		panic("cannot name any element in empty stack")
	}
	return s.fifo[len(s.fifo)-1]
}

func (s *StringStack) Parent() string {
	switch len(s.fifo) {
	case 0:
		panic("cannot name any element in empty stack")
	case 1:
		return ""
	default:
		return s.fifo[len(s.fifo)-2]
	}
}

// RawVisitor visits node and collects stats.
type RawVisitor struct {
	// Prefix names the element to focus on.
	Prefix string
	// stack keeps track of the position in the tree.
	stack StringStack
	// local keeps hierarchy information for a single element.
	local *ChildMap
	// out is a user defined processing unit for childmaps
	out chan *ChildMap
}

func NewRawVisitor(prefix string) *RawVisitor {
	return &RawVisitor{
		Prefix: prefix,
		stack:  StringStack{},
		local:  NewChildMap(),
		out:    ChildmapPrinter(),
	}
}

func (v *RawVisitor) hasPrefix() bool {
	return strings.HasPrefix(v.stack.Path(), v.Prefix)
}

func (v *RawVisitor) Visit(node interface{}) error {
	switch node := node.(type) {
	case xml.StartElement:
		v.stack.Push(node.Name.Local)
		if v.hasPrefix() {
			v.local.AddNode(v.stack.Parent(), v.stack.Name())
			for _, attr := range node.Attr {
				v.local.AddAttr(v.stack.Name(), attr.Name.Local)
			}
		}
	case xml.EndElement:
		if v.Prefix == v.stack.Path() {
			v.out <- v.local
			v.local = NewChildMap()
		}
		v.stack.Pop()
	case nil:
		close(v.out)
	default:
		return nil
	}
	return nil
}

func ChildmapPrinter() chan *ChildMap {
	ch := make(chan *ChildMap)
	go func() {
		for cm := range ch {
			b, err := json.Marshal(cm)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}
	}()
	return ch
}
