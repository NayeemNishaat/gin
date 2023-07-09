package lib

import "github.com/golodash/galidator"

var g = galidator.New()

func GetCustomizer(i interface{}) galidator.Validator {
	return g.Validator(i)
}
