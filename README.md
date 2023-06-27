# simple-json
A simple json module in Go.

## TODO

- [ ] array interface
- [ ] map interfacepackage json // import "github.com/chuqingq/simple-json"


## go doc

```
TYPES

type Json struct {
	sjson.Json
}
    Json Json

func FromBytes(data []byte) (*Json, error)
    FromBytes parse []byte to Json

func FromFile(filepath string) (*Json, error)
    FromFile parse json from filepath

func FromString(s string) (*Json, error)
    FromString parse string to Json

func FromStruct(v interface{}) *Json
    FromStruct parse struct to Json, like json.Marshal()

func New() *Json
    New new a Json

func (m *Json) Array() []Json

func (m *Json) Get(path string) *Json
    Get 按path获取Json中的内容。path是"key1.key2.key3"的形式

func (m *Json) Map(path string, def ...map[string]interface{}) map[string]interface{}

func (m *Json) Set(path string, v interface{})
    Set
    path是"key1.key2.key3"的形式，v支持string、bool、int、map[string]interface{}、[]interface{}

func (m *Json) ToBytes() []byte
    ToBytes Json to []byte

func (m *Json) ToFile(filepath string) error
    ToFile save json to filepath

func (m *Json) ToString() string
    ToString Json to string

func (m *Json) ToStruct(v interface{}) error
    ToStruct parse Json to struct, like json.Unmarshal()
```
