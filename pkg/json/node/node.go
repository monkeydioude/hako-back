package node

import (
	"bytes"
)

type Node interface {
	Process() []byte
	Add([]Node)
	GetKey() string
}

func WriteString(b *bytes.Buffer, s string) {
	b.WriteByte('"')
	b.WriteString(s)
	b.WriteByte('"')
}

func WriteKey(b *bytes.Buffer, s string) {
	WriteString(b, s)
	b.WriteByte(':')
}
