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

func (o *Object) Add(n Node) {
	o.Nodes[n.GetKey()] = n
}

func (o *Object) Bytes() []byte {
	b := bytes.NewBuffer([]byte{'{'})
	iterated := false
	for k, n := range o.Nodes {
		if iterated {
			b.WriteByte(',')
		}
		iterated = true
		WriteKey(b, k)
		b.Write(n.Bytes())
	}
	b.WriteByte('}')

	return b.Bytes()
}

func (o *Object) GetKey() string {
	return o.NodeKey
}
