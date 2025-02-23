package test_setup

import (
	"fmt"
	"testing"

	"rickover/setup"
	"rickover/test"
)

func TestActiveQueries(t *testing.T) {
	test.SetUp(t)
	defer test.TearDown(t)
	count, err := setup.GetActiveQueries()
	test.AssertNotError(t, err, "")
	test.Assert(t, count >= 1, fmt.Sprintf("Expected count >= 1, got %d", count))
}
