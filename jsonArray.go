package jsonx

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type JSONArray struct {
	list []interface{}
}

func NewJSONArrayFromList(list []interface{}) *JSONArray {
	return &JSONArray{list: list}
}
func NewJSONArray() *JSONArray {
	return &JSONArray{list: make([]interface{}, 0, 16)}
}

func (j *JSONArray) Size() int {
	return len(j.list)
}
func (j *JSONArray) IsEmpty() bool {
	return j.Size() == 0
}
func (j *JSONArray) Contains(o interface{}) bool {
	for _, v := range j.list {
		if v == o {
			return true
		}
	}
	return false
}
func (j *JSONArray) ToArray() []interface{} {
	return j.list
}
func (j *JSONArray) Add(o interface{}) {
	j.list = append(j.list, o)
}
func (j *JSONArray) FluentAdd(o interface{}) *JSONArray {
	j.Add(o)
	return j
}
func (j *JSONArray) Remove(o interface{}) {
	for i, v := range j.list {
		if v == o {
			j.list = append(j.list[:i], j.list[i+1:]...)
			break
		}
	}
}
func (j *JSONArray) FluentRemove(o interface{}) *JSONArray {
	j.Remove(o)
	return j
}
func (j *JSONArray) AddAll(o []interface{}) {
	j.list = append(j.list, o...)
}
func (j *JSONArray) FluentAddAll(o []interface{}) *JSONArray {
	j.AddAll(o)
	return j
}
func (j *JSONArray) RemoveAll(o []interface{}) {
	for _, v := range o {
		j.Remove(v)
	}
}
func (j *JSONArray) FluentRemoveAll(o []interface{}) *JSONArray {
	j.RemoveAll(o)
	return j
}
func (j *JSONArray) Clear() {
	j.list = make([]interface{}, 0, 16)
}
func (j *JSONArray) FluentClear() *JSONArray {
	j.Clear()
	return j
}
func (j *JSONArray) Set(index int, o interface{}) {
	j.list[index] = o
}
func (j *JSONArray) FluentSet(index int, o interface{}) *JSONArray {
	j.Set(index, o)
	return j
}
func (j *JSONArray) AddIndex(index int, o interface{}) {
	j.list = append(j.list[:index], append([]interface{}{o}, j.list[index:]...)...)
}

func (j *JSONArray) FluentAddIndex(index int, o interface{}) *JSONArray {
	j.AddIndex(index, o)
	return j
}
func (j *JSONArray) RemoveIndex(index int) {
	j.list = append(j.list[:index], j.list[index+1:]...)
}
func (j *JSONArray) FluentRemoveIndex(index int) *JSONArray {
	j.RemoveIndex(index)
	return j
}
func (j *JSONArray) IndexOf(o interface{}) int {
	for i, v := range j.list {
		if v == o {
			return i
		}
	}
	return -1
}
func (j *JSONArray) LastIndexOf(o interface{}) int {
	for i := len(j.list) - 1; i >= 0; i-- {
		if j.list[i] == o {
			return i
		}
	}
	return -1
}
func (j *JSONArray) Get(index int) interface{} {
	return j.list[index]
}

func (j *JSONArray) GetJSONArray(index int) *JSONArray {
	v := j.Get(index)
	switch v.(type) {
	case *JSONArray:
		return v.(*JSONArray)
	default:
		return nil
	}
}
func (j *JSONArray) GetJSONObject(index int) *JSONObject {
	v := j.Get(index)
	switch v.(type) {
	case *JSONObject:
		return v.(*JSONObject)
	case map[string]interface{}:
		return NewJSONObjectMap(v.(map[string]interface{}))
	default:
		return nil
	}
}

func (j *JSONArray) GetBoolean(index int) (bool, error) {
	v := j.Get(index)
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
func (j *JSONArray) GetBytes(index int) []byte {
	v := j.Get(index)
	switch v.(type) {
	case []byte:
		return v.([]byte)
	case string:
		return []byte(v.(string))
	default:
		return []byte{}
	}
}

func (j *JSONArray) GetByte(index int) byte {
	v := j.Get(index)
	switch v.(type) {
	case byte:
		return v.(byte)
	default:
		return 0
	}
}
func (j *JSONArray) GetIntValue(index int) int {
	return int(j.GetInt64Value(index))
}
func (j *JSONArray) GetInt64Value(index int) int64 {
	v := j.Get(index)
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
func (j *JSONArray) GetFloat64Value(index int) float64 {
	v := j.Get(index)
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

func (j *JSONArray) ForEach(call ForCallArray) {
	size := j.Size()
	for i := 0; i < size; i++ {
		if !call(i, j.GetJSONObject(i)) {
			break
		}
	}
}
func (j *JSONArray) GetString(index int) string {
	v := j.Get(index)
	switch v.(type) {
	case float64:
		return strconv.FormatInt(int64(v.(float64)), 10)
	case string:
		return v.(string)
	}
	return ""
}
func (j *JSONArray) GetUnixMilliDate(index int) time.Time {
	milli := j.GetInt64Value(index)
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}
func (j JSONArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.list)
}

func (j *JSONArray) UnmarshalJSON(b []byte) error {
	if !IsJSONArray(b) {
		return errors.New("not a json array")
	}
	return json.Unmarshal(b, &j.list)
}
