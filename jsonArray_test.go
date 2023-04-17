package jsonx

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray_ToString(t *testing.T) {
	arr := NewJSONArray()
	arr.Add(1)
	arr.Add(2)
	marshal, err := json.Marshal(arr)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
