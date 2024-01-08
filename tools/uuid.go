package tools

import "math/rand"

// Generate int64 ID
func GenId() int64 {
	return rand.Int63n(999999999999999999)
}
