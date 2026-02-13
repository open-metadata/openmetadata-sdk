package ometa

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func EncodeParams(params any) url.Values {
	if params == nil {
		return nil
	}

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	values := url.Values{}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		
		tag := t.Field(i).Tag.Get("form")
		if tag == "" || tag == "-" {
			continue
		}

		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}

		val := field
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		values.Set(tag, fmt.Sprintf("%v", val.Interface()))

	}
	return values
}