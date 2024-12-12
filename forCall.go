package jsonx

type ForCall func(key string, value interface{}) bool

type ForCallArray func(index int, value *JSONObject) bool
