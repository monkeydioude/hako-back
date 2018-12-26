package node

import "bytes"

type Object struct {
	Nodes   map[string]Node
	NodeKey string
}

func NewObject(key string) *Object {
	return &Object{
		NodeKey: key,
		Nodes:   make(map[string]Node),
	}
}

func (o *Object) Add(nodes []Node) {
	for _, n := range nodes {
		o.Nodes[n.GetKey()] = n
	}
}

func (o *Object) Process() []byte {
	b := bytes.NewBuffer([]byte{'{'})
	iterated := false
	for k, n := range o.Nodes {
		if iterated {
			b.WriteByte(',')
		}
		iterated = true
		WriteKey(b, k)
		b.Write(n.Process())
	}
	b.WriteByte('}')

	return b.Bytes()
}

func (o *Object) GetKey() string {
	return o.NodeKey
}
