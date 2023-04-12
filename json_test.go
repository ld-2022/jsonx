package jsonx

import (
	"testing"
)

func TestParseObject(t *testing.T) {
	j := ParseObject([]byte(`{"a":1}`))
	if j.GetIntValue("a") != 1 {
		t.Fail()
	}
}

func TestParseArray(t *testing.T) {
	j := ParseArray([]byte(`[1,2,3]`))
	if j.GetIntValue(0) != 1 {
		t.Fail()
	}
	if j.GetIntValue(1) != 2 {
		t.Fail()
	}
	if j.GetIntValue(2) != 3 {
		t.Fail()
	}
}

func TestToJSONString(t *testing.T) {
	j := ParseObject([]byte(`{"a":1}`))
	if j.ToJsonString() != `{"a":1}` {
		t.Fail()
	}
}

func TestToJSONBytes(t *testing.T) {
	j := ParseObject([]byte(`{"a":1}`))
	if string(j.ToJsonBytes()) != `{"a":1}` {
		t.Fail()
	}
}
