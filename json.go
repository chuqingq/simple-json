package json

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	sjson "github.com/bitly/go-simplejson"
	"github.com/chuqingq/go-util"
)

// Json comm使用的消息
type Json struct {
	sjson.Json
}

func New() *Json {
	return &Json{}
}

// Set 支持v是string、bool、int、map[string]interface{}、[]interface{}
func (m *Json) Set(path string, v interface{}) {
	m.SetPath(strings.Split(path, "."), v)
}

func (m *Json) Get(path string) *Json {
	if path == "" {
		return m
	}
	return &Json{*m.GetPath(strings.Split(path, ".")...)}
}

func (m *Json) String(path string, def ...string) string {
	return m.Get(path).MustString(def...)
}

func (m *Json) Bool(path string, def ...bool) bool {
	return m.Get(path).MustBool(def...)
}

func (m *Json) Int(path string, def ...int) int {
	return m.Get(path).MustInt(def...)
}

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
	b, err := m.EncodePretty()
	if err != nil {
		log.Printf("messagep[%v].EncodePretty() error: %v", m, err)
	}
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
	str := util.ToJson(v)
	m, _ := FromString(str)
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

/*
TODO
func New() *Message
func NewMessage(v interface{}) *Message

func (m *Message) Set(path string, val interface{})
func (m *Message) Get(path string) *Message

func (m *Message) MessageArray() ([]Message, error)
func (m *Message) MustMessageArray(msg ...[]Message) []Message

func MessageFromFile(filepath string) (*Message, error)
func (m *Message) ToFile(filepath string) error

*/
