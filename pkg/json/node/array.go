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

func (a *Array) Add(n Node) {
	a.Childs = append(a.Childs, n)
}

func (a *Array) Bytes() []byte {
	b := bytes.NewBuffer([]byte{'['})
	for i, n := range a.Childs {
		if i > 0 {
			b.WriteByte(',')
		}
		b.Write(n.Bytes())
	}
	b.WriteByte(']')

	return b.Bytes()
}

func (a *Array) GetKey() string {
	return a.NodeKey
}
