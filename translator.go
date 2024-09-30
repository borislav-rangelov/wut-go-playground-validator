package wuterr

import (
	"github.com/borislav-rangelov/wut"
	"github.com/go-playground/validator/v10"
)

type (
	Translator interface {
		Translate(lang string, err validator.FieldError) *wut.Result
		ToMap(lang string, errs []validator.FieldError) map[string]*wut.Result
	}
)

type DefaultTranslator struct {
	Prefix       string
	KeyExtractor KeyExtractor
	LangFactory  wut.LangFactory
}

func (d *DefaultTranslator) Translate(lang string, err validator.FieldError) *wut.Result {
	keys := d.KeyExtractor.FromFieldError(d.Prefix, err)
	return d.LangFactory.Lang(lang).GetFirst(keys, err)
}

func (d *DefaultTranslator) ToMap(lang string, errs []validator.FieldError) map[string]*wut.Result {
	result := make(map[string]*wut.Result)
	for _, err := range errs {
		result[getFieldName(err)] = d.Translate(lang, err)
	}
	return result
}

func getFieldName(err validator.FieldError) string {
	field := err.StructField()
	if field == "" {
		field = err.Field()
	}
	return field
}
