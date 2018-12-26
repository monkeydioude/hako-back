package json

import (
	"github.com/monkeydioude/hako-back/pkg/json/node"
)

func Array(nodes ...node.Node) *node.Array {
	a := &node.Array{}
	a.Add(nodes)
	return a
}

func Object(nodes ...node.Node) *node.Object {
	o := node.NewObject("")
	o.Add(nodes)
	return o
}

func Marshal(data interface{}) *node.Marshal {
	return &node.Marshal{
		Data: data,
	}
}

func String(value string) *node.String {
	return &node.String{
		Value: value,
	}
}

type key struct {
	key string
}

func Key(k string) *key {
	return &key{
		key: k,
	}
}

func (k *key) Array(nodes ...node.Node) *node.Array {
	a := node.NewArray(k.key)
	a.Add(nodes)
	return a
}

func (k *key) Object(nodes ...node.Node) *node.Object {
	o := node.NewObject(k.key)
	o.Add(nodes)
	return o
}

func (k *key) Marshal(data interface{}) *node.Marshal {
	m := Marshal(data)
	m.NodeKey = k.key
	return m
}

func (k *key) String(value string) *node.String {
	s := String(value)
	s.NodeKey = k.key
	return s
}
