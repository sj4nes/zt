//go:build debug

package quality

func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
