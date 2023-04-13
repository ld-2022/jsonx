package jsonx

import (
	"encoding/json"
	"errors"
	"github.com/imdario/mergo"
	"log"
	"strconv"
	"time"
)

type JSONObject struct {
	m map[string]interface{}
}

func NewJSONObjectMap(m map[string]interface{}) *JSONObject {
	return &JSONObject{m: m}
}
func NewJSONObject() *JSONObject {
	return &JSONObject{m: make(map[string]interface{}, 16)}
}

func (j *JSONObject) Size() int {
	return len(j.m)
}

func (j *JSONObject) IsEmpty() bool {
	return j.Size() == 0
}
func (j *JSONObject) ContainsKey(key string) bool {
	_, ok := j.m[key]
	return ok
}
func (j *JSONObject) ContainsValue(value interface{}) bool {
	for _, v := range j.m {
		if v == value {
			return true
		}
	}
	return false
}
func (j *JSONObject) Get(key string) interface{} {
	return j.m[key]
}
func (j *JSONObject) GetJSONObject(key string) *JSONObject {
	v := j.Get(key)
	switch v.(type) {
	case map[string]interface{}:
		return NewJSONObjectMap(v.(map[string]interface{}))
	case *JSONObject:
		return v.(*JSONObject)
	default:
		return ParseObject(ToJSONBytes(v))
	}
}

func (j *JSONObject) getJSONArray(key string) *JSONArray {
	v := j.Get(key)
	switch v.(type) {
	case *JSONArray:
		return v.(*JSONArray)
	default:
		return nil
	}
}
func (j *JSONObject) GetBoolean(key string) (bool, error) {
	v := j.Get(key)
	switch v.(type) {
	case bool:
		return v.(bool), nil
	case int:
		return v.(int) == 1, nil
	case string:
		strV := v.(string)
		if strV != "" {
			if "true" != strV && "1" != strV {
				if "false" != strV && "0" != strV {
					if "Y" != strV && "T" != strV {
						if "F" != strV && "N" != strV {
							return false, errors.New("can not cast to boolean, value : " + strV)
						} else {
							return false, nil
						}
					} else {
						return true, nil
					}
				} else {
					return false, nil
				}
			} else {
				return true, nil
			}
		} else {
			return false, nil
		}
	default:
		return false, errors.New("can not cast to boolean, value : ")
	}
}
func (j *JSONObject) GetBytes(key string) []byte {
	v := j.Get(key)
	switch v.(type) {
	case []byte:
		return v.([]byte)
	case string:
		return []byte(v.(string))
	default:
		return []byte{}
	}
}

func (j *JSONObject) GetByte(key string) byte {
	v := j.Get(key)
	switch v.(type) {
	case byte:
		return v.(byte)
	default:
		return 0
	}
}
func (j *JSONObject) GetIntValue(key string) int {
	return int(j.GetInt64Value(key))
}
func (j *JSONObject) GetInt64Value(key string) int64 {
	v := j.Get(key)
	switch v.(type) {
	case int:
		return int64(v.(int))
	case float64:
		return int64(v.(float64))
	case int64:
		return v.(int64)
	}
	return 0
}
func (j *JSONObject) GetFloat64Value(key string) float64 {
	v := j.Get(key)
	switch v.(type) {
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	case int64:
		return float64(v.(int64))
	}
	return 0
}

func (j *JSONObject) GetString(key string) string {
	v := j.Get(key)
	switch v.(type) {
	case float64:
		return strconv.FormatInt(int64(v.(float64)), 10)
	case string:
		return v.(string)
	}
	return ""
}
func (j *JSONObject) GetUnixMilliDate(key string) time.Time {
	milli := j.GetInt64Value(key)
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}
func (j *JSONObject) Put(key string, value interface{}) {
	j.m[key] = value
}
func (j *JSONObject) FluentPut(key string, value interface{}) *JSONObject {
	j.Put(key, value)
	return j
}
func (j *JSONObject) PutAll(m map[string]interface{}) {
	err := mergo.Merge(&j.m, m)
	if err != nil {
		log.Println(err)
	}
}
func (j *JSONObject) FluentPutAll(m map[string]interface{}) *JSONObject {
	j.PutAll(m)
	return j
}
func (j *JSONObject) Clear() {
	j.m = make(map[string]interface{})
}
func (j *JSONObject) FluentClear() *JSONObject {
	j.Clear()
	return j
}
func (j *JSONObject) Remove(key string) interface{} {
	v := j.Get(key)
	delete(j.m, key)
	return v
}
func (j *JSONObject) FluentRemove(key string) *JSONObject {
	j.Remove(key)
	return j
}
func (j *JSONObject) ToJsonString() string {
	return ToJSONString(j.m)
}

func (j *JSONObject) ToJsonBytes() []byte {
	return ToJSONBytes(j.m)
}

func (j JSONObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.m)
}

func (j *JSONObject) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &j.m)
	if err != nil {
		log.Println(err)
		return nil
	}
	return err
}
