package utilo

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	emptyTag       = ""
	tagKVSeparator = ":"
)

// Marshal marshals the given val and returns the map that contains the fields.
// optional=another is not implemented, and it's hard to implement and not common used.
func Marshal(val any) (map[string]map[string]any, error) {
	ret := make(map[string]map[string]any)
	tp := reflect.TypeOf(val)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	for i := 0; i < tp.NumField(); i++ {
		// todo:
	}

	return ret, nil
}

func getTag(field reflect.StructField) (string, bool) {
	tag := string(field.Tag)
	if i := strings.Index(tag, tagKVSeparator); i >= 0 {
		return strings.TrimSpace(tag[:i]), true
	}

	return strings.TrimSpace(tag), false
}

func validateOptional(field reflect.StructField, value reflect.Value) error {
	switch field.Type.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return fmt.Errorf("field %q is nil", field.Name)
		}
	case reflect.Array, reflect.Slice, reflect.Map:
		if value.IsNil() || value.Len() == 0 {
			return fmt.Errorf("field %q is empty", field.Name)
		}
	}

	return nil
}
