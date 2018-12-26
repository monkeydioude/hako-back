package json

import (
	"encoding/json"
	"testing"
)

func TestICanBuildArrayWithStrings(t *testing.T) {
	goal := `["a","b","c"]`
	trial := Array(
		String("a"),
		String("b"),
		String("c"),
	)
	res := trial.Process()

	if string(res) != goal {
		t.Fail()
	}
}

func TestICanBuildObjectWithString(t *testing.T) {
	goal := make(map[string]string)
	trial := make(map[string]string)
	json.Unmarshal([]byte(`{"a":"A","b":"B","c":"C"}`), &goal)

	object := Object(
		Key("a").String("A"),
		Key("b").String("B"),
		Key("c").String("C"),
	)

	json.Unmarshal(object.Process(), &trial)

	if trial["a"] != goal["a"] ||
		trial["b"] != goal["b"] ||
		trial["c"] != goal["c"] {
		t.Fail()
	}
}

func TestICanBuildObjectWithStringsAndArrays(t *testing.T) {
	goal := make(map[string][]map[string]string)
	trial := make(map[string][]map[string]string)
	json.Unmarshal([]byte(`{"c":[{
			"z": "Z",
			"y": "Y",
			"x" : "X"
		}]}`),
		&goal,
	)

	object := Object(
		Key("c").Array(
			Object(
				Key("z").String("Z"),
				Key("y").String("Y"),
				Key("x").String("X"),
			),
		),
	)

	json.Unmarshal(object.Process(), &trial)

	if trial["c"][0]["z"] != goal["c"][0]["z"] ||
		trial["c"][0]["y"] != goal["c"][0]["y"] ||
		trial["c"][0]["x"] != goal["c"][0]["x"] {
		t.Fail()
	}
}
