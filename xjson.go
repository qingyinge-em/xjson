package xjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Xjson struct {
	str string
}

func New() *Xjson {
	return new(Xjson)
}

func NewFromString(jstr string) (*Xjson, error) {
	j := &Xjson{str: jstr}
	if gjson.Valid(jstr) {
		return j, nil
	}
	return nil, fmt.Errorf("invalid json string")
}

func NewFromFile(file string) (*Xjson, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewFromString(string(data))
}

func (j *Xjson) Assign(str string) {
	j.str = str
}

func (j *Xjson) Set(path string, val any) {
	j.str, _ = sjson.Set(j.str, path, val)
}

func (j *Xjson) Delete(path string) {
	j.str, _ = sjson.Delete(j.str, path)
}

func (j *Xjson) Reset() {
	j.str = ""
}

func (j *Xjson) String() string {
	return j.str
}

func (j *Xjson) IndentString() string {
	var b bytes.Buffer
	err := json.Indent(&b, []byte(j.str), "", "  ")
	if err != nil {
		return fmt.Sprintf("json.Indent err:%v:%s", err, j.str)
	}
	return b.String()
}

func (j *Xjson) CompactString() string {
	var b bytes.Buffer
	err := json.Compact(&b, []byte(j.str))
	if err != nil {
		return fmt.Sprintf("json.Compact err:%v:%s", err, j.str)
	}
	return b.String()
}

func (j *Xjson) Exist(path string) bool {
	return gjson.Get(j.str, path).Exists()
}

func (j *Xjson) GetString(path string) string {
	return gjson.Get(j.str, path).String()
}

func (j *Xjson) GetInt(path string) int {
	return int(gjson.Get(j.str, path).Int())
}

func (j *Xjson) GetBool(path string) bool {
	return gjson.Get(j.str, path).Bool()
}

func (j *Xjson) GetArraySize(path string) int {
	return int(gjson.Get(j.str, path+".#").Int())
}

func (j *Xjson) Unmarshal(path string, val any) error {
	if path == "" {
		return jsoniter.UnmarshalFromString(j.str, val)
	} else {
		return jsoniter.UnmarshalFromString(gjson.Get(j.str, path).String(), val)
	}
}

// when array, key is "0","1",...
func (j *Xjson) ForEach(path string,
	cb func(key string, value string) bool) {

	gjson.Get(j.str, path).ForEach(
		func(key, value gjson.Result) bool {
			return cb(key.String(), value.String())
		},
	)
}
