package node

import "encoding/json"

type Marshal struct {
	Data    interface{}
	NodeKey string
}

func (m *Marshal) Add(n Node) {}

func (m *Marshal) GetKey() string {
	return m.NodeKey
}

func (m *Marshal) Bytes() []byte {
	res, err := json.Marshal(m.Data)
	if err != nil {
		return nil
	}
	return res
}
