package uniqsym

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var Optional = []validation.Rule{
	validation.Length(1, 512),
	validation.Match(regexp.MustCompile(`^` + regex + `$`)),
}

var Required = append(Optional, validation.Required)
