package symbol

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var Optional = []validation.Rule{
	validation.Length(1, 128),
	validation.Match(regexp.MustCompile(`^` + Chars + `*$`)),
}

var Required = append(Optional, validation.Required)
