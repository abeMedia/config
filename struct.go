package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type configStruct struct {
	c interface{}
}

func (c *configStruct) ReadFile(p string) error {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return fmt.Errorf("config: %s", err)
	}

	switch path.Ext(p) {
	case ".json":
		return c.ReadJSON(b)
	case ".yaml", ".yml":
		return c.ReadYAML(b)
	}
	return ErrUnsupportedFormat
}

func (c *configStruct) ReadJSON(b []byte) error {
	return json.Unmarshal(b, c.c)
}

func (c *configStruct) ReadYAML(b []byte) error {
	return yaml.Unmarshal(b, c.c)
}

func (c *configStruct) String() string {
	b, err := json.Marshal(c.c)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (c *configStruct) Get(path string) (interface{}, error) {
	v, err := c.getValue(path)
	if err != nil {
		return nil, err
	}
	return v.Interface(), nil
}

func (c *configStruct) Set(path string, value interface{}) (err error) {
	v, err := c.getValue(path)
	if err != nil {
		return err
	}
	val := reflect.ValueOf(value)
	if v.Kind() != val.Kind() {
		return fmt.Errorf("config: type mismatch: cannot set `%s` (type %T) to `%s` (type %T)", path, v.Interface(), value, value)
	}
	v.Set(val)
	return nil
}

func (c *configStruct) getValue(path string) (reflect.Value, error) {
	keySlice := strings.Split(path, ".")
	v := reflect.ValueOf(c.c)
	//iterate through field names ,ignore the first name as it might be the current instance name
	for _, key := range keySlice {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		// we only accept structs
		if v.Kind() != reflect.Struct {
			return v, fmt.Errorf("config: field not found: %s", path)
		}

		v = v.FieldByName(key)
		if !v.IsValid() {
			return v, fmt.Errorf("config: field not found: %s", path)
		}
	}
	return v, nil
}
