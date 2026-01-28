package main

import (
	"github.com/mendersoftware/mender-server/tests/runner/tests"

	_ "github.com/mendersoftware/mender-server/tests/runner/tests/basic_operations"
)

func main() {
	tests.Main()
}
