package immutable_map

type Node struct {
	b     byte
	nodes Nodes
	value interface{}
}

func (a *Node) Insert(path []byte, value interface{}) *Node {
	return &Node{
		b:     a.b,
		nodes: a.nodes.Insert(path[1:], value),
	}
}

func (a *Node) Contains(path []byte) bool {
	if len(path) == 0 {
		return true
	}
	return a.nodes.Contains(path)
}

func (a *Node) Get(path []byte) (interface{}, bool) {
	return a.nodes.Get(path)
}

func newNode(path []byte, value interface{}) *Node {
	if len(path) == 1 {
		return &Node{
			b:     path[0],
			nodes: Nodes{},
			value: value,
		}
	}
	return &Node{
		b:     path[0],
		nodes: Nodes{}.Insert(path[1:], value),
	}
}

type Nodes []*Node

func (a Nodes) Insert(path []byte, value interface{}) Nodes {
	if len(path) == 0 {
		return a
	}
	exists, index := findPos(a, path[0])
	clone := dup(a)
	if exists {
		clone[index] = clone[index].Insert(path, value)
		return clone
	}
	clone = append(clone[:index], append(Nodes{newNode(path, value)}, clone[index:]...)...)
	return clone
}

func (a Nodes) Contains(path []byte) bool {
	exists, index := findPos(a, path[0])
	if !exists {
		return false
	}
	return a[index].Contains(path[1:])
}

func (a Nodes) Get(path []byte) (interface{}, bool) {
	exists, index := findPos(a, path[0])
	if !exists {
		return nil, false
	}
	if len(path) == 1 {
		return a[index].value, true
	}
	return a[index].Get(path[1:])
}

func dup(nodes []*Node) []*Node {
	out := make([]*Node, len(nodes))
	for i, v := range nodes {
		out[i] = &*v
	}
	return out
}

func findPos(nodes []*Node, b byte) (exists bool, pos int) {
	for i, v := range nodes {
		if b <= v.b {
			return true, i
		}
		if b > v.b {
			continue
		}
	}
	return false, len(nodes)
}

type Map struct {
	nodes Nodes
}

func New() *Map {
	return &Map{}
}

func (a Map) Contains(path []byte) bool {
	if len(path) == 0 {
		return false
	}
	return a.nodes.Contains(path)
}

func (a *Map) Insert(path []byte, value interface{}) *Map {
	return &Map{
		nodes: a.nodes.Insert(path, value),
	}
}

func (a Map) Get(path []byte) (interface{}, bool) {
	return a.nodes.Get(path)
}