package lib

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidateStrongPass(field validator.FieldLevel) bool {
	// strings.Contains(field.Field().String(), "")
	// match, _ := regexp.MatchString("^[a-aA-Z]$", field.Field().String())

	str := field.Field().String()
	letters := 0
	number := false
	upper := false
	special := false

	for _, c := range str {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
		}
	}

	if letters > 7 && number && upper && special {
		return true
	}
	return false
}

// ^(?=.*([A-Z]){1,})(?=.*[!@#$&*]{1,})(?=.*[0-9]{1,})(?=.*[a-z]{1,}).{8,100}$ // Remark: Password Regex

// Segment:
// ^(?=.u)(?=U)..$
// Point: .u matches any+u and forgets it and set the cursor to the beginning/previous cursor if any (0)
// Point: U matches U and forgets it and set the cursor to the previous position which is the beginning (0)
// Point: .. matches any two from the beginning because those positive look around doesn't consume any characters. After getting a match they are simply discarded and cursor was set to the previous position (0)

// Important:
// ^(?=u)(?=U)..$ will never match because at first this will try to find u at the 0 position then set cursor to 0. After that it will look for a U at the 0 position which is impossible because u and U cannot exist at the same 0 position.
