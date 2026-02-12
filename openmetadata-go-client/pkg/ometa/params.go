package ometa

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// We'll take the params struct and convert it to url.Values for query parameters
// This is a simple implementation that uses reflection to read struct tags and values
func EncodeParams(params any) url.Values {
	if params == nil {
		return nil
	}

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem() // Dereference pointer
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
		
		// Read the `form` tag. Example: `form:"fields,omitempty"` → "fields"
		tag := t.Field(i).Tag.Get("form")
		if tag == "" || tag == "-" {
			continue
		}

		// Strip the ",omitempty" part — we just need the parameter name
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