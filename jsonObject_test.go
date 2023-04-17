package jsonx

import (
	"fmt"
	"testing"
)

func TestNewJSONObject(t *testing.T) {
	j := NewJSONObject()
	if j == nil {
		t.Error("NewJSONObject() failed")
	}
}

func TestJSONObject_ForEach(t *testing.T) {
	j := NewJSONObject()
	j.Put("a", 1)
	j.Put("b", 2)
	j.Put("c", 3)
	j.ForEach(func(key string, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
