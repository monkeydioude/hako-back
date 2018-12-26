package node

import "bytes"

type String struct {
	NodeKey string
	Value   string
}

func NewString(k, v string) *String {
	return &String{
		NodeKey: k,
		Value:   v,
	}
}

func (s *String) Add(n []Node) {}

func (s *String) Process() []byte {
	b := &bytes.Buffer{}
	WriteString(b, s.Value)
	return b.Bytes()
}

func (s *String) GetKey() string {
	return s.NodeKey
}
