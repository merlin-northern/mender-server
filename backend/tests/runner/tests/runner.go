package tests

import (
	"strconv"
	"testing"
)

type TestSettings struct {
	ServerURL string
	Username  string
	Password  string
}

type MainTestFunc func(t *testing.T, settings TestSettings) error

var TestCases = make(map[string]MainTestFunc)

func AddTestCase(name string, main MainTestFunc) {
	i := 0
	for {
		if _, defined := TestCases[name]; defined {
			i++
			name = name + "_" + strconv.Itoa(i)
		} else {
			break
		}
	}
	TestCases[name] = main
}
