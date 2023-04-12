package jsonx

import "testing"

func TestNewJSONObject(t *testing.T) {
	j := NewJSONObject()
	if j == nil {
		t.Error("NewJSONObject() failed")
	}
}
