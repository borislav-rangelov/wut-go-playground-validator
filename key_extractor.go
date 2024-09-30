package wuterr

import "github.com/go-playground/validator/v10"

type KeyExtractor interface {
	FromFieldError(prefix string, err validator.FieldError) []string
}

type DefaultKeyExtractor struct{}

func (d *DefaultKeyExtractor) FromFieldError(prefix string, err validator.FieldError) []string {
	result := make([]string, 0)
	prefixKey := prefix + "."

	if err.StructNamespace() != "" {
		// validation.User.Email.required
		result = append(result, prefixKey+err.StructNamespace()+"."+err.Tag())
	}
	if err.StructField() != "" {
		// validation.Email.required
		result = append(result, prefixKey+err.StructField()+"."+err.Tag())
	}
	// validation.required
	result = append(result, prefixKey+err.Tag())
	return result
}
