package validation

import "reflect"

// isValid do a simple check for strings and integers
func isValid(field interface{}) bool {
	value := reflect.ValueOf(field)
	valid := false

	switch value.Kind() {
	case reflect.String:
		if value.String() != "" {
			valid = true
		}
	case reflect.Float64:
		if value.Float() >= 1 {
			valid = true
		}
	}

	return valid
}
