package validator

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

type ValidationEnvelope struct {
	validator *Validator
	prefix    string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Envelope(prefix string) *ValidationEnvelope {
	return &ValidationEnvelope{
		validator: v,
		prefix:    prefix,
	}
}

func (v *Validator) ArrayEnvelope(arrayKey string, index int) *ValidationEnvelope {
	prefix := fmt.Sprintf("%s.%d", arrayKey, index)
	return &ValidationEnvelope{
		validator: v,
		prefix:    prefix,
	}
}

func (env *ValidationEnvelope) Check(ok bool, key, message string) {
	if !ok {
		fullKey := key
		if env.prefix != "" {
			fullKey = strings.Join([]string{env.prefix, key}, ".")
		}
		env.validator.AddError(fullKey, message)
	}
}

func (env *ValidationEnvelope) Envelope(prefix string) *ValidationEnvelope {
	newPrefix := prefix
	if env.prefix != "" {
		newPrefix = strings.Join([]string{env.prefix, prefix}, ".")
	}
	return &ValidationEnvelope{
		validator: env.validator,
		prefix:    newPrefix,
	}
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
