package json

import (
	"encoding/json"
	"os"
	"strings"

	sjson "github.com/bitly/go-simplejson"
)

// Json Json
type Json struct {
	sjson.Json
}

// New new a Json
func New() *Json {
	return &Json{}
}

// Set path是"key1.key2.key3"的形式，v支持string、bool、int、map[string]interface{}、[]interface{}
func (m *Json) Set(path string, v interface{}) {
	m.SetPath(strings.Split(path, "."), v)
}

// Get 按path获取Json中的内容。path是"key1.key2.key3"的形式
func (m *Json) Get(path string) *Json {
	if path == "" {
		return m
	}
	return &Json{*m.GetPath(strings.Split(path, ".")...)}
}

// func (j *Json) MustString(def...string)
// func (j *Json) MustBool(def...bool)
// func (j *Json) MustInt(def...int)
// func (j *Json) MustInt64(def...int64)
// func (j *Json) MustFloat64(def...float64)

func (m *Json) Array() []Json {
	arr, err := m.Json.Array()
	if arr == nil || err != nil {
		return nil
	}
	marray := make([]Json, len(arr))
	for i, a := range arr {
		marray[i].SetPath([]string{}, a)
	}
	return marray
}

func (m *Json) Map(path string, def ...map[string]interface{}) map[string]interface{} {
	return m.Get(path).MustMap(def...)
}

// Unmarshal 把m解析到v上。类似json.Unmarshal()
func (m *Json) Unmarshal(v interface{}) error {
	b := m.ToBytes()
	return json.Unmarshal(b, v)
}

// ToBytes Message转成[]byte
func (m *Json) ToBytes() []byte {
	b, _ := m.EncodePretty()
	return b
}

// ToString Message转成string
func (m *Json) ToString() string {
	return string(m.ToBytes())
}

// FromBytes 字节数组转成Message
func FromBytes(data []byte) (*Json, error) {
	m, err := sjson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return &Json{*m}, nil
}

// FromString 字符串转成Message
func FromString(s string) (*Json, error) {
	m, err := FromBytes([]byte(s))
	return m, err
}

// FromStruct 类似json.Marshal()
func FromStruct(v interface{}) *Json {
	b, _ := json.Marshal(v)
	m, _ := FromBytes(b)
	return m
}

// FromFile 从filepath读取Message
func FromFile(filepath string) (*Json, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	m, err := FromBytes(b)
	return m, err
}

// ToFile 把Message保存到filepath
func (m *Json) ToFile(filepath string) error {
	b, err := m.EncodePretty()
	if err != nil {
		return err
	}
	const defaultFileMode = 0644
	return os.WriteFile(filepath, b, defaultFileMode)
}
