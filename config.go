package config

import (
	"errors"
	"reflect"
)

// Config ...
type Config interface {
	ReadFile(p string) error
	ReadJSON(b []byte) error
	ReadYAML(b []byte) error
	String() string
	Get(path string) (interface{}, error)
	Set(path string, value interface{}) (err error)
}

// New ...
func New(c interface{}) (Config, error) {
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return nil, ErrInvalidType
	}
	switch v.Elem().Kind() {
	case reflect.Struct:
		return &configStruct{c}, nil
	case reflect.Map:
		panic("Not implemented")
	}
	return nil, ErrInvalidType
}

var (
	// ErrUnsupportedFormat is returned when reading an unsupported file format
	ErrUnsupportedFormat = errors.New("config: unsupported file format")

	// ErrInvalidType is returned when reading an unsupported file format
	ErrInvalidType = errors.New("config: invalid type: only accepts struct pointers and maps ")
)
