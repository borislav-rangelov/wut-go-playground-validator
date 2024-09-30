package wuterr

import "github.com/borislav-rangelov/wut"

type Options interface {
	Prefix(prefix string) Options
	KeyExtractor(e KeyExtractor) Options
	LangFactory(lf wut.LangFactory) Options
	Build() Translator
	AsDefault()
}

type options struct {
	prefix       string
	keyExtractor KeyExtractor
	langFactory  wut.LangFactory
}

var defaultTranslator Translator

func GetDefaultTranslator() Translator {
	return defaultTranslator
}

type defaultLangFactory struct{}

func (d *defaultLangFactory) Lang(code string) wut.LangSource {
	return wut.Lang(code)
}

func Setup() Options {
	return &options{
		prefix:       "validation",
		keyExtractor: nil,
		langFactory:  nil,
	}
}

func (o options) Prefix(prefix string) Options {
	if prefix != "" {
		o.prefix = prefix
	} else {
		o.prefix = "validation"
	}
	return o
}

func (o options) KeyExtractor(e KeyExtractor) Options {
	o.keyExtractor = e
	return o
}

func (o options) LangFactory(lf wut.LangFactory) Options {
	o.langFactory = lf
	return o
}

func (o options) Build() Translator {
	lf := o.langFactory
	if lf == nil {
		lf = &defaultLangFactory{}
	}
	ke := o.keyExtractor
	if ke == nil {
		ke = &DefaultKeyExtractor{}
	}
	return &DefaultTranslator{
		Prefix:       o.prefix,
		KeyExtractor: ke,
		LangFactory:  lf,
	}
}

func (o options) AsDefault() {
	defaultTranslator = o.Build()
}
