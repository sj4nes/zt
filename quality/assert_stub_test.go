//go:build !debug

package quality

import "testing"

func TestAssert(t *testing.T) {
	Assert(true, "This should not panic")
	Assert(false, "This should not panic in a release build")
}
