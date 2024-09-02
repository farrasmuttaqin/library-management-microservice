package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
)

// IsEmptyStruct check anything in golang is empty
func IsEmptyStruct(value interface{}) bool {
	v := reflect.ValueOf(value)

	// Check if the value is valid
	if !v.IsValid() {
		// If the value is not valid, consider it empty and return true
		return true
	}

	// Handle different kinds of types
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		// Check if the length of array, map, slice, or string is zero
		return v.Len() == 0
	case reflect.Bool:
		// Check if boolean value is false
		return !v.Bool()
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		// Check if value is nil
		return v.IsNil()
	case reflect.Complex64, reflect.Complex128:
		// Check if complex value is zero
		return v.Complex() == 0
	case reflect.Float32, reflect.Float64:
		// Check if float value is zero
		return v.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Check if integer value is zero
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		// Check if unsigned integer value is zero
		return v.Uint() == 0
	case reflect.Struct:
		// Check if the struct is empty
		return IsEmptyStructDetail(v)
	default:
		// For other types, consider it not empty
		return false
	}
}
func IsEmptyStructDetail(v reflect.Value) bool {
	// Iterate through each field in the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// Check if the field is exported
		if field.CanInterface() {
			// Get the default value for the field type
			zero := reflect.Zero(field.Type())
			// Compare the field value with the default value
			if !reflect.DeepEqual(field.Interface(), zero.Interface()) {
				// If at least one field is not empty, return false
				return false
			}
		}
	}
	// If all fields are empty or unexported, return true
	return true
}

func ExtractPayload(req *http.Request) string {
	if req.Body == nil {
		return "No Payload"
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "Error: Unable to read request body"
	}

	// Set body back for further use
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return string(body)
}
