package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	validate          = "validate"
	required          = "required"
	min               = "min"
	max               = "max"
	email             = "email"
	regex             = "regex"
	emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

// ValidationError represents a single validation error
type ValidatorError struct {
	Field   string
	Message string
}

// Custom DataTypes
type (
	// ValidationErrors holds multiple validation errors
	ValidationErrors []ValidatorError
)

// ValidationErrors impl Error Interface by calling error function
func (ve ValidationErrors) Error() string {
	var errMsgs []string

	for _, errVal := range ve {
		errMsgs = append(errMsgs, fmt.Sprintf("%s : %s", errVal.Field, errVal.Message))
	}
	return strings.Join(errMsgs, "; ")
}

// Validator handles validation logic
type Validator struct {
	errors ValidationErrors
}

func New() *Validator {
	return &Validator{}
}

// /// Validate performs  validation on the provided struct
func (ve *Validator) Validate(s interface{}) error {
	ve.errors = ValidationErrors{}
	rVal := reflect.ValueOf(s)
	if rVal.Kind() != reflect.Pointer {
		return fmt.Errorf("validation requires a struct pointer input")
	}

	// Get the type of the struct  ex:Person struct
	var structVal reflect.Value
	structVal = rVal.Elem()
	if structVal.Kind() != reflect.Struct {
		return fmt.Errorf("refOut must be a pointer struct !")
	}
	ve.validateFields(structVal)
	// process
	// validate fields
	if len(ve.errors) > 0 {
		return ve.errors // calling Error fn
	}

	return nil
}

func (v *Validator) validateFields(structVal reflect.Value) {

	//  get type
	structTyp := structVal.Type()

	for i := 0; i < structVal.NumField(); i++ {
		currentField := structTyp.Field(i)
		currentFieldVal := structVal.Field(i)

		// Get validation rules from struct tag `validate:"required,min=2,max=50"`
		tagVal := currentField.Tag.Get(validate)
		if tagVal == "" {
			continue
		}

		rules := strings.Split(tagVal, ",")
		for _, rule := range rules {
			if err := v.applyValidationRule(rule, currentFieldVal); err != nil {
				v.errors = append(v.errors, ValidatorError{
					Field:   currentField.Name,
					Message: err.Error(),
				})
			}

		}

	}

}

// applyValidationRule  rules => current rule for current field elem and its valuefor element itself
func (v *Validator) applyValidationRule(rule string, currentFieldVal reflect.Value) error {

	parts := strings.Split(rule, "=")
	ruleName := strings.Trim(parts[0], " ")
	var ruleValue string

	// handle require
	if len(parts) > 1 {
		ruleValue = strings.Trim(parts[1], " ")
	}

	switch ruleName {
	case required:
		if currentFieldVal.IsZero() {
			return fmt.Errorf("field is required")
		}
	case min:
		return v.validateMin(ruleValue, currentFieldVal)
	case email:
		if !v.isMatchedRegex(currentFieldVal.String(), emailRegexPattern) {
			return fmt.Errorf("invalid email format")
		}
	case regex:
		if !v.isMatchedRegex(currentFieldVal.String(), ruleValue) {
			return fmt.Errorf("invalid email format")
		}
	}

	return nil
}

func (v *Validator) validateMin(minvalue string, currentFieldVal reflect.Value) error {

	min, err := strconv.Atoi(minvalue)
	if err != nil {
		return fmt.Errorf("invalid min value")
	}

	switch currentFieldVal.Kind() {
	case reflect.String:
		if len(currentFieldVal.String()) < min {
			return fmt.Errorf("length must be at least %d", min)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if currentFieldVal.Int() < int64(min) {
			return fmt.Errorf("value must be at least %d", min)
		}

	}
	return nil

}
func (v *Validator) validateMax(maxvalue string, currentFieldVal reflect.Value) error {

	max, err := strconv.Atoi(maxvalue)
	if err != nil {
		return fmt.Errorf("invalid min value")
	}

	switch currentFieldVal.Kind() {
	case reflect.String:
		if len(currentFieldVal.String()) > max {
			return fmt.Errorf("length must be at most %d", max)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if currentFieldVal.Int() > int64(max) {
			return fmt.Errorf("length must be at most %d", max)
		}
	}
	return nil

}

func (v *Validator) isMatchedRegex(value, pattern string) bool {

	matched, _ := regexp.MatchString(pattern, value)
	return matched
}
