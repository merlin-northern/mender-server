package runner

import (
	"fmt"

	"github.com/mendersoftware/mender-server/tests/runner/tests"
)

func init() {
	tests.AddTestCase("test_tiers", mainTestTiers)
}

func mainTestTiers() error {
	fmt.Printf("test_tiers starting\n")
	return nil
}
