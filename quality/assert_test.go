//go:build debug

package quality

import "testing"

func TestAssert(t *testing.T) {
	// Test with condition true
	Assert(true, "This should not panic")

	// Test with condition false
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but did not get one in debug build")
		}
	}()
	Assert(false, "This should panic")
}
