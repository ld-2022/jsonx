package jsonx

type JsonAware interface {
	ToJsonString() string
}
