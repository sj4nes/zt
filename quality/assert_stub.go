//go:build !debug

package quality

func Assert(condition bool, message string) {
	// This function is a no-op in release builds
}
