package node

import "bytes"

type Array struct {
	Childs  []Node
	NodeKey string
}

func NewArray(key string) *Array {
	return &Array{
		NodeKey: key,
	}
}

func (a *Array) Add(nodes []Node) {
	for _, n := range nodes {
		a.Childs = append(a.Childs, n)
	}
}

func (a *Array) Process() []byte {
	b := bytes.NewBuffer([]byte{'['})
	for i, n := range a.Childs {
		if i > 0 {
			b.WriteByte(',')
		}
		b.Write(n.Process())
	}
	b.WriteByte(']')

	return b.Bytes()
}

func (a *Array) GetKey() string {
	return a.NodeKey
}
