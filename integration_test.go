package wuterr

import (
	"github.com/borislav-rangelov/wut"
	"github.com/go-playground/validator/v10"
	"testing"
)

type TestUser struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

func TestName(t *testing.T) {

	err := wut.Setup().AddFiles("testdata/example_en.toml").AsDefault()
	if err != nil {
		t.Fatal(err)
	}

	Setup().AsDefault()

	validate := validator.New(validator.WithRequiredStructEnabled())

	errs := validate.Struct(&TestUser{
		Username: "",
		Email:    "something",
		Password: "short",
	})

	translated := GetDefaultTranslator().ToMap("en", errs.(validator.ValidationErrors))

	if len(translated) != 3 {
		t.Errorf("translated length should be 3, got %d", len(translated))
	}

	expectValue(t, translated, "Username", "Required")
	expectValue(t, translated, "Email", "'something' is not a valid email")
	expectValue(t, translated, "Password", "Password must be between 8 and 32 characters")
}

func expectValue(t *testing.T, translated map[string]*wut.Result, key string, expected string) {
	if result, ok := translated[key]; ok {
		if result.Txt != expected {
			t.Errorf("Translated %s should have been '%s', got %s", key, expected, result.Txt)
		}
	} else {
		t.Errorf("%s error should not be nil", key)
	}
}
