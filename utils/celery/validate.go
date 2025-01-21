package celery

import (
	"errors"
	"fmt"
	"reflect"
)

// TaskParams is an interface for task-specific parameters.
type TaskParams interface {
	Validate() error // Each task parameter must implement a validation method.
}

type BaseParams struct{}

// ValidateStruct validates the fields of a struct based on tags.
func ValidateStruct[T any](param T) error {
	v := reflect.ValueOf(param)

	// Ensure the input is a struct or pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return errors.New("expected a struct or pointer to struct")
	}

	// Iterate over the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		// Check for custom "validate" tag
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}

		switch tag {
		case "required":
			if isZeroValue(value) {
				return fmt.Errorf("field %s is required", field.Name)
			}
		case "positive":
			if value.Kind() == reflect.Int || value.Kind() == reflect.Int64 || value.Kind() == reflect.Int32 {
				if value.Int() <= 0 {
					return fmt.Errorf("field %s must be positive", field.Name)
				}
			}
		// Add more cases as needed for additional validation rules
		default:
			return fmt.Errorf("unknown validation tag: %s", tag)
		}
	}

	return nil
}

// isZeroValue checks if a reflect.Value is its type's zero value.
func isZeroValue(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
