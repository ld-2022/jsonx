package jsonx

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func ParseObject(bytes []byte) *JSONObject {
	// 判断是否是json字符串
	if !IsJSONObject(bytes) {
		return nil
	}
	j := NewJSONObject()
	if err := json.Unmarshal(bytes, &j.m); err != nil {
		return nil
	}
	return j
}

func ParseArray(bytes []byte) *JSONArray {
	// 判断是否是json字符串
	if !IsJSONArray(bytes) {
		return nil
	}
	j := NewJSONArray()
	if err := json.Unmarshal(bytes, &j.list); err != nil {
		return nil
	}
	return j
}

func ToJSONString(object interface{}) string {
	return string(ToJSONBytes(object))
}

func ToJSONBytes(object interface{}) []byte {
	marshal, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
		return nil
	}
	return marshal
}

func skipWS(s string) string {
	if len(s) == 0 || s[0] > 0x20 {
		// Fast path.
		return s
	}
	return skipWSSlow(s)
}

func skipWSSlow(s string) string {
	if len(s) == 0 || s[0] != 0x20 && s[0] != 0x0A && s[0] != 0x09 && s[0] != 0x0D {
		return s
	}
	for i := 1; i < len(s); i++ {
		if s[i] != 0x20 && s[i] != 0x0A && s[i] != 0x09 && s[i] != 0x0D {
			return s[i:]
		}
	}
	return ""
}
func parseRawString(s string) (string, string, error) {
	n := strings.IndexByte(s, '"')
	if n < 0 {
		return s, "", fmt.Errorf(`missing closing '"'`)
	}
	if n == 0 || s[n-1] != '\\' {
		// Fast path. No escaped ".
		return s[:n], s[n+1:], nil
	}

	// Slow path - possible escaped " found.
	ss := s
	for {
		i := n - 1
		for i > 0 && s[i-1] == '\\' {
			i--
		}
		if uint(n-i)%2 == 0 {
			return ss[:len(ss)-len(s)+n], s[n+1:], nil
		}
		s = s[n+1:]

		n = strings.IndexByte(s, '"')
		if n < 0 {
			return ss, "", fmt.Errorf(`missing closing '"'`)
		}
		if n == 0 || s[n-1] != '\\' {
			return ss[:len(ss)-len(s)+n], s[n+1:], nil
		}
	}
}

// IsJSONObject 判断是否是json对象
func IsJSONObject(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}
	return bytes[0] == '{' && bytes[len(bytes)-1] == '}'
}

// IsJSONArray 判断是否是json数组
func IsJSONArray(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}
	return bytes[0] == '[' && bytes[len(bytes)-1] == ']'
}
