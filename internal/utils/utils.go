package utils

import "fmt"

func Assert(b bool, text string, args ...interface{}) {
	if b {
		panic(fmt.Sprintf(text, args...))
	}
}

func AssertErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
