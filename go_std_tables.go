// Code generated by scripts/gen_go_std_tables.go; DO NOT EDIT.

// Generated from Go versions [go1.23.0].

package main

var runtimeAndDeps = map[string]bool{
	"internal/abi":              true, // go1.23.0
	"internal/bytealg":          true, // go1.23.0
	"internal/byteorder":        true, // go1.23.0
	"internal/chacha8rand":      true, // go1.23.0
	"internal/coverage/rtcov":   true, // go1.23.0
	"internal/cpu":              true, // go1.23.0
	"internal/goarch":           true, // go1.23.0
	"internal/godebugs":         true, // go1.23.0
	"internal/goexperiment":     true, // go1.23.0
	"internal/goos":             true, // go1.23.0
	"internal/profilerecord":    true, // go1.23.0
	"internal/runtime/atomic":   true, // go1.23.0
	"internal/runtime/exithook": true, // go1.23.0
	"internal/runtime/syscall":  true, // go1.23.0
	"internal/stringslite":      true, // go1.23.0
	"runtime":                   true, // go1.23.0
	"runtime/internal/math":     true, // go1.23.0
	"runtime/internal/sys":      true, // go1.23.0
	"unsafe":                    true, // go1.23.0
}

var runtimeLinknamed = []string{
	"arena",                          // go1.23.0
	"crypto/internal/boring",         // go1.23.0
	"crypto/internal/boring/bcache",  // go1.23.0
	"crypto/internal/boring/fipstls", // go1.23.0
	"crypto/x509/internal/macos",     // go1.23.0
	"internal/coverage/cfile",        // go1.23.0
	"internal/godebug",               // go1.23.0
	"internal/poll",                  // go1.23.0
	"internal/reflectlite",           // go1.23.0
	"internal/syscall/unix",          // go1.23.0
	"internal/syscall/windows",       // go1.23.0
	"internal/weak",                  // go1.23.0
	"maps",                           // go1.23.0
	"os",                             // go1.23.0
	"os/signal",                      // go1.23.0
	"plugin",                         // go1.23.0
	"reflect",                        // go1.23.0
	"runtime/debug",                  // go1.23.0
	"runtime/metrics",                // go1.23.0
	"runtime/pprof",                  // go1.23.0
	"runtime/trace",                  // go1.23.0
	"sync",                           // go1.23.0
	"sync/atomic",                    // go1.23.0
	"syscall",                        // go1.23.0
	"syscall/js",                     // go1.23.0
	"time",                           // go1.23.0
	"unique",                         // go1.23.0
	// The net package linknames to the runtime, not the other way around.
	// TODO: support this automatically via our script.
	"net",
}

var compilerIntrinsics = map[string]map[string]bool{
	"internal/runtime/atomic": {
		"And":             true, // go1.23.0
		"And32":           true, // go1.23.0
		"And64":           true, // go1.23.0
		"And8":            true, // go1.23.0
		"Anduintptr":      true, // go1.23.0
		"Cas":             true, // go1.23.0
		"Cas64":           true, // go1.23.0
		"CasRel":          true, // go1.23.0
		"Casint32":        true, // go1.23.0
		"Casint64":        true, // go1.23.0
		"Casp1":           true, // go1.23.0
		"Casuintptr":      true, // go1.23.0
		"Load":            true, // go1.23.0
		"Load64":          true, // go1.23.0
		"Load8":           true, // go1.23.0
		"LoadAcq":         true, // go1.23.0
		"LoadAcq64":       true, // go1.23.0
		"LoadAcquintptr":  true, // go1.23.0
		"Loadint32":       true, // go1.23.0
		"Loadint64":       true, // go1.23.0
		"Loadp":           true, // go1.23.0
		"Loaduint":        true, // go1.23.0
		"Loaduintptr":     true, // go1.23.0
		"Or":              true, // go1.23.0
		"Or32":            true, // go1.23.0
		"Or64":            true, // go1.23.0
		"Or8":             true, // go1.23.0
		"Oruintptr":       true, // go1.23.0
		"Store":           true, // go1.23.0
		"Store64":         true, // go1.23.0
		"Store8":          true, // go1.23.0
		"StoreRel":        true, // go1.23.0
		"StoreRel64":      true, // go1.23.0
		"StoreReluintptr": true, // go1.23.0
		"Storeint32":      true, // go1.23.0
		"Storeint64":      true, // go1.23.0
		"StorepNoWB":      true, // go1.23.0
		"Storeuintptr":    true, // go1.23.0
		"Xadd":            true, // go1.23.0
		"Xadd64":          true, // go1.23.0
		"Xaddint32":       true, // go1.23.0
		"Xaddint64":       true, // go1.23.0
		"Xadduintptr":     true, // go1.23.0
		"Xchg":            true, // go1.23.0
		"Xchg64":          true, // go1.23.0
		"Xchgint32":       true, // go1.23.0
		"Xchgint64":       true, // go1.23.0
		"Xchguintptr":     true, // go1.23.0
	},
	"math": {
		"Abs":         true, // go1.23.0
		"Ceil":        true, // go1.23.0
		"Copysign":    true, // go1.23.0
		"FMA":         true, // go1.23.0
		"Floor":       true, // go1.23.0
		"Round":       true, // go1.23.0
		"RoundToEven": true, // go1.23.0
		"Trunc":       true, // go1.23.0
		"sqrt":        true, // go1.23.0
	},
	"math/big": {
		"mulWW": true, // go1.23.0
	},
	"math/bits": {
		"Add":             true, // go1.23.0
		"Add64":           true, // go1.23.0
		"Div":             true, // go1.23.0
		"Div64":           true, // go1.23.0
		"Len":             true, // go1.23.0
		"Len16":           true, // go1.23.0
		"Len32":           true, // go1.23.0
		"Len64":           true, // go1.23.0
		"Len8":            true, // go1.23.0
		"Mul":             true, // go1.23.0
		"Mul64":           true, // go1.23.0
		"OnesCount":       true, // go1.23.0
		"OnesCount16":     true, // go1.23.0
		"OnesCount32":     true, // go1.23.0
		"OnesCount64":     true, // go1.23.0
		"OnesCount8":      true, // go1.23.0
		"Reverse":         true, // go1.23.0
		"Reverse16":       true, // go1.23.0
		"Reverse32":       true, // go1.23.0
		"Reverse64":       true, // go1.23.0
		"Reverse8":        true, // go1.23.0
		"ReverseBytes16":  true, // go1.23.0
		"ReverseBytes32":  true, // go1.23.0
		"ReverseBytes64":  true, // go1.23.0
		"RotateLeft":      true, // go1.23.0
		"RotateLeft16":    true, // go1.23.0
		"RotateLeft32":    true, // go1.23.0
		"RotateLeft64":    true, // go1.23.0
		"RotateLeft8":     true, // go1.23.0
		"Sub":             true, // go1.23.0
		"Sub64":           true, // go1.23.0
		"TrailingZeros16": true, // go1.23.0
		"TrailingZeros32": true, // go1.23.0
		"TrailingZeros64": true, // go1.23.0
		"TrailingZeros8":  true, // go1.23.0
	},
	"runtime": {
		"publicationBarrier": true, // go1.23.0
	},
	"runtime/internal/math": {
		"MulUintptr": true, // go1.23.0
	},
	"runtime/internal/sys": {
		"Bswap32":          true, // go1.23.0
		"Bswap64":          true, // go1.23.0
		"Len64":            true, // go1.23.0
		"Len8":             true, // go1.23.0
		"OnesCount64":      true, // go1.23.0
		"Prefetch":         true, // go1.23.0
		"PrefetchStreamed": true, // go1.23.0
		"TrailingZeros32":  true, // go1.23.0
		"TrailingZeros64":  true, // go1.23.0
		"TrailingZeros8":   true, // go1.23.0
	},
	"sync": {
		"runtime_LoadAcquintptr":  true, // go1.23.0
		"runtime_StoreReluintptr": true, // go1.23.0
	},
	"sync/atomic": {
		"AddInt32":              true, // go1.23.0
		"AddInt64":              true, // go1.23.0
		"AddUint32":             true, // go1.23.0
		"AddUint64":             true, // go1.23.0
		"AddUintptr":            true, // go1.23.0
		"AndInt32":              true, // go1.23.0
		"AndInt64":              true, // go1.23.0
		"AndUint32":             true, // go1.23.0
		"AndUint64":             true, // go1.23.0
		"AndUintptr":            true, // go1.23.0
		"CompareAndSwapInt32":   true, // go1.23.0
		"CompareAndSwapInt64":   true, // go1.23.0
		"CompareAndSwapUint32":  true, // go1.23.0
		"CompareAndSwapUint64":  true, // go1.23.0
		"CompareAndSwapUintptr": true, // go1.23.0
		"LoadInt32":             true, // go1.23.0
		"LoadInt64":             true, // go1.23.0
		"LoadPointer":           true, // go1.23.0
		"LoadUint32":            true, // go1.23.0
		"LoadUint64":            true, // go1.23.0
		"LoadUintptr":           true, // go1.23.0
		"OrInt32":               true, // go1.23.0
		"OrInt64":               true, // go1.23.0
		"OrUint32":              true, // go1.23.0
		"OrUint64":              true, // go1.23.0
		"OrUintptr":             true, // go1.23.0
		"StoreInt32":            true, // go1.23.0
		"StoreInt64":            true, // go1.23.0
		"StoreUint32":           true, // go1.23.0
		"StoreUint64":           true, // go1.23.0
		"StoreUintptr":          true, // go1.23.0
		"SwapInt32":             true, // go1.23.0
		"SwapInt64":             true, // go1.23.0
		"SwapUint32":            true, // go1.23.0
		"SwapUint64":            true, // go1.23.0
		"SwapUintptr":           true, // go1.23.0
	},
}

var reflectSkipPkg = map[string]bool{
	"fmt": true,
}
