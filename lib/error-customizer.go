package lib

import "github.com/golodash/galidator"

var g = galidator.New()

func GetCustomizer(i any) galidator.Validator {
	return g.Validator(i)
}

// Note: interface{} == *void == any. It's called the empty interface and is implemented by all types, which means you can put anything in the Msg field.
// Important: The interface{} type (or any with Go 1.18+), the empty interface is the interface that has no methods.
